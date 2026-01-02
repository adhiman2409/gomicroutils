package gomicroutils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
	"os"
	"strings"
)

// EncryptString encrypts a plaintext string using the provided key string
// The key can be any length - it will be hashed to create a 32-byte AES key
// Returns base64url encoded ciphertext (URL-safe: uses - and _ instead of + and /)
func EncryptString(plaintext, key string) (string, error) {
	if plaintext == "" {
		return "", errors.New("plaintext cannot be empty")
	}
	if key == "" {
		return "", errors.New("key cannot be empty")
	}

	// Create a 32-byte key from the provided key string using SHA-256
	hashedKey := sha256.Sum256([]byte(key))

	// Create AES cipher block
	block, err := aes.NewCipher(hashedKey[:])
	if err != nil {
		return "", err
	}

	// Convert plaintext to bytes
	plaintextBytes := []byte(plaintext)

	// Create GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Create a nonce (number used once)
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Encrypt and append nonce to the beginning
	ciphertext := gcm.Seal(nonce, nonce, plaintextBytes, nil)

	// Encode to base64url (URL-safe format)
	base64Str := base64.StdEncoding.EncodeToString(ciphertext)
	base64url := strings.ReplaceAll(base64Str, "+", "-")
	base64url = strings.ReplaceAll(base64url, "/", "_")
	base64url = strings.TrimRight(base64url, "=")

	return base64url, nil
}

// DecryptString decrypts a base64url encoded ciphertext using the provided key string
// The key must be the same string used for encryption
// Returns the original plaintext
func DecryptString(ciphertext, skey string) (string, error) {
	if ciphertext == "" {
		return "", errors.New("ciphertext cannot be empty")
	}
	key := skey
	if key == "" {
		key = os.Getenv("ENCRYPTION_KEY")
	}

	// Convert base64url to standard base64
	// Replace URL-safe characters back to standard base64 characters
	base64Str := strings.ReplaceAll(ciphertext, "-", "+")
	base64Str = strings.ReplaceAll(base64Str, "_", "/")

	// Add padding if needed
	// base64 strings should be a multiple of 4 characters
	if padding := len(base64Str) % 4; padding > 0 {
		base64Str += strings.Repeat("=", 4-padding)
	}

	// Decode base64
	ciphertextBytes, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return "", errors.New("invalid ciphertext format")
	}

	// Create a 32-byte key from the provided key string using SHA-256
	hashedKey := sha256.Sum256([]byte(key))

	// Create AES cipher block
	block, err := aes.NewCipher(hashedKey[:])
	if err != nil {
		return "", err
	}

	// Create GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Check minimum size
	nonceSize := gcm.NonceSize()
	if len(ciphertextBytes) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	// Extract nonce and ciphertext
	nonce, ciphertextBytes := ciphertextBytes[:nonceSize], ciphertextBytes[nonceSize:]

	// Decrypt
	plaintextBytes, err := gcm.Open(nil, nonce, ciphertextBytes, nil)
	if err != nil {
		return "", errors.New("decryption failed - invalid key or corrupted data")
	}

	return string(plaintextBytes), nil
}
