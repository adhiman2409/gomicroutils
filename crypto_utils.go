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
)

// EncryptString encrypts a plaintext string using the provided key string
// The key can be any length - it will be hashed to create a 32-byte AES key
// Returns base64 encoded ciphertext
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

	// Encode to base64 for easy storage/transmission
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// DecryptString decrypts a base64 encoded ciphertext using the provided key string
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

	// Decode base64
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
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
