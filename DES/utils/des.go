package utils

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
	"errors"
)

// DESEncrypt encrypts the given plain text using DES algorithm with the provided key.
func DESEncrypt(plainText, key string) (string, error) {
	block, err := des.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	if len(plainText)%block.BlockSize() != 0 {
		return "", errors.New("plain text is not a multiple of the block size")
	}

	ciphertext := make([]byte, len(plainText))
	mode := cipher.NewCBCEncrypter(block, make([]byte, block.BlockSize()))
	mode.CryptBlocks(ciphertext, []byte(plainText))

	return hex.EncodeToString(ciphertext), nil
}

// DESDecrypt decrypts the given cipher text using DES algorithm with the provided key.
func DESDecrypt(cipherText, key string) (string, error) {
	block, err := des.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	ciphertext, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	if len(ciphertext)%block.BlockSize() != 0 {
		return "", errors.New("cipher text is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, make([]byte, block.BlockSize()))
	mode.CryptBlocks(ciphertext, ciphertext)

	return string(ciphertext), nil
}

func PadText(text string) string {
	blockSize := des.BlockSize
	padding := blockSize - len(text)%blockSize
	padText := text + string(paddingBytes(padding))
	return padText
}

//UnpadText removes PKCS7 padding from the input text.

func UnpadText(text string) string {
	padding := int(text[len(text)-1])
	return text[:len(text)-padding]
}

func paddingBytes(n int) []byte {
	b := byte(n)
	return bytes.Repeat([]byte{b}, n)
}
