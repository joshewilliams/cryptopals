package aes

import (
	"crypto/aes"
)

// ECB function to perform AES-ECB-128 encryption/decryption as needed
func ECB(data, key []byte) []byte {
	cipher, _ := aes.NewCipher([]byte(key))
	out := make([]byte, len(data))
	size := 16

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		cipher.Decrypt(out[bs:be], data[bs:be])
	}

	return out
}
