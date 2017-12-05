package brute

import (
	"cryptopals/set1/charscore"
	"cryptopals/set1/xor"
	"encoding/hex"
	"strings"
)

// Xor function for bruteforcing xor'd data, single byte key
func Xor(src string) (string, int) {
	high := 0
	var out string
	for i := 0; i <= 255; i++ {
		key := hex.EncodeToString([]byte(string(i)))
		result := xor.ToASCII(src, key, true, false)
		score := charscore.TotalScore(result)
		if score > high {
			high = score
			out = result
		}
	}
	return out, high
}

// Repeat function for bruteforcing xor'd data with a repeating xor key (between 2 and 40 bytes in length)
func Repeat(src string) string {
	//bSrc := []byte(src)
	var low, finalKeySize int
	for keysize := 2; keysize <= 40; keysize++ {
		tmp1, tmp2, tmp3, tmp4 := src[:keysize], src[keysize:keysize*2], src[keysize*2:keysize*3], src[keysize*3:keysize*4]
		result1 := charscore.Hamming(tmp1, tmp2)
		result2 := charscore.Hamming(tmp2, tmp3)
		result3 := charscore.Hamming(tmp3, tmp4)
		result := (result1 + result2 + result3) / 3
		if keysize == 2 {
			low = result
			finalKeySize = keysize
		} else {
			if result < low {
				low = result
				finalKeySize = keysize
			}
		}
	}
	block := make([]string, len(src)/finalKeySize)
	for i := 0; i < len(src)/finalKeySize; i++ {
		block[i] = src[finalKeySize*i : finalKeySize*i+1]
	}
	block2 := make([]string, len(block))
	for i := 0; i < len(block[i]); i++ {
		var tmp byte
		for j := 0; j < len(block); j++ {
			tmp += block[j][i]
		}
		block2[i] = string(tmp)
	}
	out := make([]string, len(block2))
	for i := range block2 {
		out[i], _ = Xor(block2[i])
	}
	finalOut := strings.Join(out, "")
	return finalOut
}
