package hrms

// Whole-document encryption for payroll collections.
//
// Each stored payroll document keeps a small set of plaintext top-level keys (the
// ones used as MongoDB query/join filters) and packs everything else — all salary
// amounts, tax, bank details and PII — into a single encrypted blob field (encKey).
// Anyone querying Mongo directly sees only the query keys plus opaque ciphertext.
//
// Reads are dual-mode: a legacy plaintext document (no encKey field) decodes
// unchanged, so the encryption-aware build can be deployed before the data is
// migrated. Encryption uses AES-256-GCM with the key from the ENCRYPTION_KEY env var.

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
	"os"

	"go.mongodb.org/mongo-driver/bson"
)

// encKey is the reserved field holding the encrypted blob. Its presence also marks a
// document as already-encrypted (used for dual-mode reads and idempotent migration).
const encKey = "__enc"

// Plaintext top-level keys per document type — every other top-level field is
// encrypted. These are exactly the keys used in MongoDB filters, so queries keep
// working. Anything not listed here (including nested PAN/UAN/Aadhaar) is encrypted.
var (
	plaintextSalaryStructure = set(
		"_id", "employee_id", "employee_name", "employment_status", "employee_type",
	)
	// PayrollSalaryConfig backs both employee-salary_configs and
	// organization-salary_configs, so the allowlist is the union of both query sets.
	plaintextSalaryConfig = set(
		"_id", "employee_id", "employee_type", "payroll_config_name",
		"state", "country", "is_default_for_employee_type", "is_fulltime_employee_salary_config",
	)
	plaintextPayrollMaster = set(
		"_id", "employee_id", "financial_year", "assessment_year", "employment_status",
	)
	plaintextProcessingRequest = set(
		"_id", "request_id", "month", "financial_year", "request_status",
		"approver_id", "verifier_id", "requested_by_id", "country", "state",
	)
)

func set(keys ...string) map[string]bool {
	m := make(map[string]bool, len(keys))
	for _, k := range keys {
		m[k] = true
	}
	return m
}

// encryptionKey returns the AES key, failing fast if it is not configured so payroll
// data is never silently persisted in plaintext.
func encryptionKey() (string, error) {
	k := os.Getenv("ENCRYPTION_KEY")
	if k == "" {
		return "", errors.New("ENCRYPTION_KEY is not set; refusing to (de)serialize payroll data")
	}
	return k, nil
}

// marshalEncrypted marshals v to BSON, then moves every top-level field that is not
// in plaintext into an encrypted blob stored under encKey.
func marshalEncrypted(v interface{}, plaintext map[string]bool) ([]byte, error) {
	key, err := encryptionKey()
	if err != nil {
		return nil, err
	}
	raw, err := bson.Marshal(v)
	if err != nil {
		return nil, err
	}
	var doc bson.D
	if err := bson.Unmarshal(raw, &doc); err != nil {
		return nil, err
	}

	out := bson.D{}
	sensitive := bson.D{}
	for _, e := range doc {
		if e.Key == encKey {
			continue // never nest a stale blob
		}
		if plaintext[e.Key] {
			out = append(out, e)
		} else {
			sensitive = append(sensitive, e)
		}
	}

	if len(sensitive) > 0 {
		payload, err := bson.Marshal(sensitive)
		if err != nil {
			return nil, err
		}
		ciphertext, err := aesEncrypt(payload, key)
		if err != nil {
			return nil, err
		}
		out = append(out, bson.E{Key: encKey, Value: ciphertext})
	}
	return bson.Marshal(out)
}

// unmarshalEncrypted decodes BSON into v, decrypting the blob first. It is dual-mode:
// a legacy plaintext document (no encKey field) is decoded directly.
func unmarshalEncrypted(data []byte, v interface{}) error {
	var doc bson.D
	if err := bson.Unmarshal(data, &doc); err != nil {
		return err
	}

	merged := bson.D{}
	var blob string
	hasBlob := false
	for _, e := range doc {
		if e.Key == encKey {
			if s, ok := e.Value.(string); ok {
				blob, hasBlob = s, true
			}
			continue
		}
		merged = append(merged, e)
	}

	if !hasBlob {
		return bson.Unmarshal(data, v) // legacy plaintext
	}

	key, err := encryptionKey()
	if err != nil {
		return err
	}
	payload, err := aesDecrypt(blob, key)
	if err != nil {
		return err
	}
	var sensitive bson.D
	if err := bson.Unmarshal(payload, &sensitive); err != nil {
		return err
	}
	merged = append(merged, sensitive...)

	raw, err := bson.Marshal(merged)
	if err != nil {
		return err
	}
	return bson.Unmarshal(raw, v)
}

// aesEncrypt encrypts plain with AES-256-GCM (key derived via SHA-256) and returns
// URL-safe base64. Self-contained to avoid an import cycle with the parent package.
func aesEncrypt(plain []byte, key string) (string, error) {
	gcm, err := newGCM(key)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	ct := gcm.Seal(nonce, nonce, plain, nil)
	return base64.RawURLEncoding.EncodeToString(ct), nil
}

// aesDecrypt reverses aesEncrypt.
func aesDecrypt(enc, key string) ([]byte, error) {
	ctBytes, err := base64.RawURLEncoding.DecodeString(enc)
	if err != nil {
		return nil, errors.New("invalid ciphertext format")
	}
	gcm, err := newGCM(key)
	if err != nil {
		return nil, err
	}
	ns := gcm.NonceSize()
	if len(ctBytes) < ns {
		return nil, errors.New("ciphertext too short")
	}
	nonce, body := ctBytes[:ns], ctBytes[ns:]
	plain, err := gcm.Open(nil, nonce, body, nil)
	if err != nil {
		return nil, errors.New("decryption failed - invalid key or corrupted data")
	}
	return plain, nil
}

func newGCM(key string) (cipher.AEAD, error) {
	hashed := sha256.Sum256([]byte(key))
	block, err := aes.NewCipher(hashed[:])
	if err != nil {
		return nil, err
	}
	return cipher.NewGCM(block)
}

// ---- BSON codecs on the payroll document structs ----
//
// The type-alias in each method drops the custom methods so the default codec is used
// for the underlying struct, avoiding infinite recursion.

func (s SalaryStructureDetails) MarshalBSON() ([]byte, error) {
	type alias SalaryStructureDetails
	return marshalEncrypted(alias(s), plaintextSalaryStructure)
}

func (s *SalaryStructureDetails) UnmarshalBSON(data []byte) error {
	type alias SalaryStructureDetails
	return unmarshalEncrypted(data, (*alias)(s))
}

func (c PayrollSalaryConfig) MarshalBSON() ([]byte, error) {
	type alias PayrollSalaryConfig
	return marshalEncrypted(alias(c), plaintextSalaryConfig)
}

func (c *PayrollSalaryConfig) UnmarshalBSON(data []byte) error {
	type alias PayrollSalaryConfig
	return unmarshalEncrypted(data, (*alias)(c))
}

func (m EmployeePayrollMaster) MarshalBSON() ([]byte, error) {
	type alias EmployeePayrollMaster
	return marshalEncrypted(alias(m), plaintextPayrollMaster)
}

func (m *EmployeePayrollMaster) UnmarshalBSON(data []byte) error {
	type alias EmployeePayrollMaster
	return unmarshalEncrypted(data, (*alias)(m))
}

func (r SalaryProcessingRequest) MarshalBSON() ([]byte, error) {
	type alias SalaryProcessingRequest
	return marshalEncrypted(alias(r), plaintextProcessingRequest)
}

func (r *SalaryProcessingRequest) UnmarshalBSON(data []byte) error {
	type alias SalaryProcessingRequest
	return unmarshalEncrypted(data, (*alias)(r))
}
