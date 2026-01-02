package main

import (
	"fmt"
	"log"

	"github.com/adhiman2409/gomicroutils"
)

func main() {
	// Original text to encrypt
	plaintext := "This is my secret message!"

	// Encryption key (can be any string)
	key := "my-super-secret-key-123"

	fmt.Println("Original text:", plaintext)
	fmt.Println("Encryption key:", key)
	fmt.Println()

	// Encrypt the string
	encrypted, err := gomicroutils.EncryptString(plaintext, key)
	if err != nil {
		log.Fatal("Encryption error:", err)
	}
	fmt.Println("Encrypted (base64):", encrypted)
	fmt.Println()

	// Decrypt the string
	decrypted, err := gomicroutils.DecryptString(encrypted, key)
	if err != nil {
		log.Fatal("Decryption error:", err)
	}
	fmt.Println("Decrypted text:", decrypted)
	fmt.Println()

	// Verify
	if plaintext == decrypted {
		fmt.Println("✓ Success! Original and decrypted text match.")
	} else {
		fmt.Println("✗ Error! Texts don't match.")
	}

	// Example with wrong key
	fmt.Println("\n--- Testing with wrong key ---")
	wrongKey := "wrong-key"
	_, err = gomicroutils.DecryptString(encrypted, wrongKey)
	if err != nil {
		fmt.Println("✓ Correctly failed with wrong key:", err)
	}
}
