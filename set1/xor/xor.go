package xor

import "encoding/hex"

// Xor function to return XOR'd byte array when given two byte arrays (Single byte key or fixed length)
func Xor(src, key []byte) []byte {
	out := make([]byte, len(src), len(src))
	if len(key) <= 2 {
		for i := range src {
			out[i] = src[i] ^ key[0]
		}
	} else {
		for i := range src {
			out[i] = src[i] ^ key[i]
		}
	}
	return out
}

// ToHex function takes input strings, XORs them and returns the hex encoded result
func ToHex(src, key string) string {
	bSrc, _ := hex.DecodeString(src)
	bKey, _ := hex.DecodeString(key)
	sHex := hex.EncodeToString(Xor(bSrc, bKey))
	return sHex
}

// ToASCII function takes input strings, XORs them and returns the ASCII result
func ToASCII(src, key string) string {
	out, _ := hex.DecodeString(ToHex(src, key))
	sAsc := string(out[:])
	return sAsc
}
