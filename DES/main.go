package main

import (
	"fmt"
	"test/utils"
)

func main() {
	var key string
	fmt.Print("Enter the key (8 characters): ")
	fmt.Scanln(&key)

	// Ensure the key length is as required for DES
	if len(key) != 8 {
		fmt.Println("Key length is not correct. Must be 8 characters.")
		return
	}

	//plainText := "Cybersecurity"
	fmt.Print("Enter the plain text: ")
	var plainText string
	fmt.Scanln(&plainText)

	// Add padding to the plain text to make its length a multiple of 8 bytes
	plainText = utils.PadText(plainText)

	encrypted, err := utils.DESEncrypt(plainText, key)
	if err != nil {
		fmt.Println("Encryption error:", err)
		return
	}
	fmt.Println("Encrypted:", encrypted)

	decrypted, err := utils.DESDecrypt(encrypted, key)
	if err != nil {
		fmt.Println("Decryption error:", err)
		return
	}

	// Remove padding after decryption
	decrypted = utils.UnpadText(decrypted)

	fmt.Println("Decrypted:", decrypted)
}
