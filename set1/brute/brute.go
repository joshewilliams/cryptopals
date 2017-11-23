package brute

import (
	"cryptopals/set1/charscore"
	"cryptopals/set1/xor"
	"encoding/hex"
)

// Xor function for bruteforcing xor'd data, single byte key
func Xor(src string) (string, int) {
	high := 0
	var out string
	for i := 0; i <= 255; i++ {
		key := hex.EncodeToString([]byte(string(i)))
		result := xor.ToASCII(src, key)
		score := charscore.TotalScore(result)
		if score > high {
			high = score
			out = result
		}
	}
	return out, high
}
