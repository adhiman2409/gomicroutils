package hrms

import (
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestGetEmployeeComplianceInfoCollection_ReturnsNamedCollection(t *testing.T) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("mongo.Connect: %v", err)
	}
	t.Cleanup(func() {
		_ = client.Disconnect(context.Background())
	})

	got := GetEmployeeComplianceInfoCollection(client, "test.domain")

	if got.Name() != PAYROLL_COLLECTION_EMPLOYEE_COMPLIANCE_INFO {
		t.Errorf("collection name: got %q want %q", got.Name(), PAYROLL_COLLECTION_EMPLOYEE_COMPLIANCE_INFO)
	}
}

func sampleEmployeeComplianceInfo() EmployeeComplianceInfo {
	return EmployeeComplianceInfo{
		ID:                     primitive.NewObjectID(),
		EmployeeID:             "PT24003",
		AadhaarLinked:          true,
		ESICInsuranceNumber:    "3412345678901234",
		IsPersonWithDisability: false,
		UpdatedAt:              time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC),
	}
}

func TestEmployeeComplianceInfoEncryptsSensitiveKeepsQueryKeys(t *testing.T) {
	t.Setenv("ENCRYPTION_KEY", testKey)
	info := sampleEmployeeComplianceInfo()

	stored := storedForm(t, info)

	// Query keys stay plaintext.
	if stored["employee_id"] != "PT24003" {
		t.Errorf("employee_id should be plaintext, got %v", stored["employee_id"])
	}
	if stored["aadhaar_linked"] != true {
		t.Errorf("aadhaar_linked should be plaintext, got %v", stored["aadhaar_linked"])
	}
	// The sensitive fields are not present in plaintext at the top level. Disability status
	// is sensitive personal data: it must live only inside the encrypted blob, never as a
	// readable top-level key.
	if _, ok := stored["esic_insurance_number"]; ok {
		t.Error("esic_insurance_number must not be stored in plaintext")
	}
	if _, ok := stored["is_person_with_disability"]; ok {
		t.Errorf("is_person_with_disability must not be stored in plaintext, got %v", stored["is_person_with_disability"])
	}
	// The encrypted blob is present.
	blob, ok := stored[encKey].(string)
	if !ok || blob == "" {
		t.Fatalf("expected encrypted blob under %q, got %v", encKey, stored[encKey])
	}
	// The raw stored bytes must not leak the sensitive value.
	raw, _ := bson.Marshal(info)
	if bytesContains(raw, info.ESICInsuranceNumber) {
		t.Errorf("sensitive value %q leaked into stored document", info.ESICInsuranceNumber)
	}
}

func TestEmployeeComplianceInfoRoundTrip(t *testing.T) {
	t.Setenv("ENCRYPTION_KEY", testKey)
	info := sampleEmployeeComplianceInfo()

	raw, err := bson.Marshal(info)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var got EmployeeComplianceInfo
	if err := bson.Unmarshal(raw, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if got.EmployeeID != info.EmployeeID {
		t.Errorf("employee_id: got %q want %q", got.EmployeeID, info.EmployeeID)
	}
	if got.ESICInsuranceNumber != info.ESICInsuranceNumber {
		t.Errorf("esic_insurance_number: got %q want %q", got.ESICInsuranceNumber, info.ESICInsuranceNumber)
	}
	if got.AadhaarLinked != info.AadhaarLinked {
		t.Errorf("aadhaar_linked: got %v want %v", got.AadhaarLinked, info.AadhaarLinked)
	}
	if got.IsPersonWithDisability != info.IsPersonWithDisability {
		t.Errorf("is_person_with_disability: got %v want %v", got.IsPersonWithDisability, info.IsPersonWithDisability)
	}
	if !got.UpdatedAt.Equal(info.UpdatedAt) {
		t.Errorf("updated_at: got %v want %v", got.UpdatedAt, info.UpdatedAt)
	}
}

// TestEmployeeComplianceInfoDisabilityStatusRoundTripsThroughBlob covers the true case
// explicitly (the sample uses false, which is also bool's zero value).
func TestEmployeeComplianceInfoDisabilityStatusRoundTripsThroughBlob(t *testing.T) {
	t.Setenv("ENCRYPTION_KEY", testKey)
	info := sampleEmployeeComplianceInfo()
	info.IsPersonWithDisability = true

	if _, ok := storedForm(t, info)["is_person_with_disability"]; ok {
		t.Fatal("is_person_with_disability must not be stored in plaintext")
	}
	raw, err := bson.Marshal(info)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var got EmployeeComplianceInfo
	if err := bson.Unmarshal(raw, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if !got.IsPersonWithDisability {
		t.Error("is_person_with_disability = false after round trip, want true")
	}
}

// TestEmployeeComplianceInfoReadsLegacyPlaintextDisabilityStatus: documents written before
// v1.19.9-compliance-phase1.3 carry is_person_with_disability as a plaintext top-level key
// alongside the blob. They must still decode correctly (the next upsert rewrites them with the
// field inside the blob).
func TestEmployeeComplianceInfoReadsLegacyPlaintextDisabilityStatus(t *testing.T) {
	t.Setenv("ENCRYPTION_KEY", testKey)
	info := sampleEmployeeComplianceInfo()
	info.IsPersonWithDisability = true

	// Reproduce the pre-phase1.3 stored form exactly: the same codec, with the old allowlist
	// that kept is_person_with_disability as a plaintext top-level key (so it is NOT inside
	// the blob).
	type alias EmployeeComplianceInfo
	legacyAllowlist := set("_id", "employee_id", "aadhaar_linked", "is_person_with_disability", "updated_at")
	legacy, err := marshalEncrypted(alias(info), legacyAllowlist)
	if err != nil {
		t.Fatalf("marshal legacy: %v", err)
	}
	var legacyDoc bson.M
	if err := bson.Unmarshal(legacy, &legacyDoc); err != nil {
		t.Fatalf("inspect legacy: %v", err)
	}
	if legacyDoc["is_person_with_disability"] != true {
		t.Fatalf("legacy fixture should carry is_person_with_disability in plaintext, got %v", legacyDoc)
	}

	var got EmployeeComplianceInfo
	if err := bson.Unmarshal(legacy, &got); err != nil {
		t.Fatalf("unmarshal legacy: %v", err)
	}
	if !got.IsPersonWithDisability || got.ESICInsuranceNumber != info.ESICInsuranceNumber {
		t.Errorf("legacy decode = %+v, want IsPersonWithDisability=true and the ESIC number intact", got)
	}
}

func TestEmployeeComplianceInfoFailsFastWithoutKey(t *testing.T) {
	t.Setenv("ENCRYPTION_KEY", "")
	info := sampleEmployeeComplianceInfo()
	if _, err := bson.Marshal(info); err == nil {
		t.Error("expected marshal to fail when ENCRYPTION_KEY is unset")
	}
}
