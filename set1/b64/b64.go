package b64

import "encoding/base64"
import "encoding/hex"

// Encode Function to convert byte slices to base64
func Encode(data string) string {
	src, _ := hex.DecodeString(data)
	dst := []byte(src)
	sEnc := base64.StdEncoding.EncodeToString(dst)
	return sEnc
}

// Decode function to convert base64 to byte slices
func Decode(data string) []byte {
	sDec, _ := base64.StdEncoding.DecodeString(data)
	return sDec
}

// ToASCII function to convert b64 encoded string to ascii
func ToASCII(data string) string {
	out := Decode(data)
	sAsc := string(out[:])
	return sAsc
}

// ToHex function to convert b64 encoded string to hex encoded string
func ToHex(data string) string {
	src := Decode(data)
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	sHex := string(dst[:])
	return sHex
}
