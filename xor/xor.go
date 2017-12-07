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

// Repeat function return XOR'd byte array when given two byte arrays (repeating XOR key)
func Repeat(src, key []byte) []byte {
	out := make([]byte, len(src), len(src))
	keylen := len(key)
	j := 0
	for i := range src {
		if j == keylen {
			j = 0
		}
		out[i] = src[i] ^ key[j]
		j++
	}
	return out
}

// ToHex function takes input strings, XORs them and returns the hex encoded result
func ToHex(src string, key string, hexa bool, repeat bool) string {
	var bSrc []byte
	var bKey []byte
	if hexa == true {
		bSrc, _ = hex.DecodeString(src)
		bKey, _ = hex.DecodeString(key)
		if len(bKey) == 0 {
			bKey = []byte(key)
		}
	} else {
		bSrc = []byte(src)
		bKey = []byte(key)
	}
	var sHex string
	if repeat == true {
		sHex = hex.EncodeToString(Repeat(bSrc, bKey))
	} else {
		sHex = hex.EncodeToString(Xor(bSrc, bKey))
	}
	return sHex
}

// ToASCII function takes input strings, XORs them and returns the ASCII result
func ToASCII(src string, key string, hexa bool, repeat bool) string {
	out, _ := hex.DecodeString(ToHex(src, key, hexa, repeat))
	sAsc := string(out[:])
	return sAsc
}
