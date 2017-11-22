package xor

// Xor function to return byte arrays when given two byte arrays
func Xor(src, dst []byte) []byte {
	var out []byte
	for i := range src {
		out[i] = src[i] ^ dst[i]
	}
	return out
}

// ToHex function takes input strings, XORs them and returns the hex encoded result
func ToHex(src, dst string) string {

}
