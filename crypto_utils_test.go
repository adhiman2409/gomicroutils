package gomicroutils

import (
	"testing"
)

func TestEncryptDecryptString(t *testing.T) {
	tests := []struct {
		name      string
		plaintext string
		key       string
	}{
		{
			name:      "simple text with simple key",
			plaintext: "Hello, World!",
			key:       "my-secret-key",
		},
		{
			name:      "long text",
			plaintext: "This is a much longer text that includes multiple sentences. It should still encrypt and decrypt properly.",
			key:       "another-secret",
		},
		{
			name:      "special characters",
			plaintext: "Test with special chars: !@#$%^&*()_+-=[]{}|;:',.<>?",
			key:       "complex-key-123!@#",
		},
		{
			name:      "unicode characters",
			plaintext: "Hello 世界 🌍",
			key:       "unicode-key",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Encrypt
			encrypted, err := EncryptString(tt.plaintext, tt.key)
			if err != nil {
				t.Fatalf("EncryptString() error = %v", err)
			}

			if encrypted == "" {
				t.Fatal("EncryptString() returned empty string")
			}

			if encrypted == tt.plaintext {
				t.Fatal("EncryptString() returned plaintext unchanged")
			}

			// Decrypt
			decrypted, err := DecryptString(encrypted, tt.key)
			if err != nil {
				t.Fatalf("DecryptString() error = %v", err)
			}

			if decrypted != tt.plaintext {
				t.Errorf("DecryptString() = %v, want %v", decrypted, tt.plaintext)
			}
		})
	}
}

func TestEncryptStringErrors(t *testing.T) {
	tests := []struct {
		name      string
		plaintext string
		key       string
		wantErr   bool
	}{
		{
			name:      "empty plaintext",
			plaintext: "",
			key:       "key",
			wantErr:   true,
		},
		{
			name:      "empty key",
			plaintext: "text",
			key:       "",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := EncryptString(tt.plaintext, tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncryptString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDecryptStringErrors(t *testing.T) {
	tests := []struct {
		name       string
		ciphertext string
		key        string
		wantErr    bool
	}{
		{
			name:       "empty ciphertext",
			ciphertext: "",
			key:        "key",
			wantErr:    true,
		},
		{
			name:       "empty key",
			ciphertext: "text",
			key:        "",
			wantErr:    true,
		},
		{
			name:       "invalid base64",
			ciphertext: "not-valid-base64!!!",
			key:        "key",
			wantErr:    true,
		},
		{
			name:       "wrong key",
			ciphertext: "dmFsaWRiYXNlNjQ=",
			key:        "wrong-key",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := DecryptString(tt.ciphertext, tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecryptString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDifferentKeysProduceDifferentResults(t *testing.T) {
	plaintext := "secret message"
	key1 := "key1"
	key2 := "key2"

	encrypted1, err := EncryptString(plaintext, key1)
	if err != nil {
		t.Fatalf("EncryptString() error = %v", err)
	}

	encrypted2, err := EncryptString(plaintext, key2)
	if err != nil {
		t.Fatalf("EncryptString() error = %v", err)
	}

	if encrypted1 == encrypted2 {
		t.Error("Different keys should produce different ciphertexts")
	}

	// Verify decryption with wrong key fails
	_, err = DecryptString(encrypted1, key2)
	if err == nil {
		t.Error("DecryptString() should fail with wrong key")
	}
}
