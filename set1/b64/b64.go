package b64

import "encoding/base64"

// Encode Function to convert byte slices to base64
func Encode(data []byte) string {
	sEnc := base64.StdEncoding.EncodeToString(data)
	return sEnc
}

// Decode function to convert base64 to byte slices
func Decode(data string) []byte {
	sDec, _ := base64.StdEncoding.DecodeString(data)
	return sDec
}
