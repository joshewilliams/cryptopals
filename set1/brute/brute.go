package brute

import (
	"cryptopals/set1/charscore"
	"cryptopals/set1/xor"
	// "encoding/hex"
	// "fmt"
	//"strings"
)

// Xor function for bruteforcing xor'd data, single byte key
func Xor(src string) (string, int) {
	high := 0
	var out string
	for i := 0; i <= 255; i++ {
		key := string(i)
		result := xor.ToASCII(src, key, false, false)
		score := charscore.TotalScore(result)
		if score > high {
			high = score
			out = key
		}
	}
	return out, high
}

// Repeat function for bruteforcing xor'd data with a repeating xor key (between 2 and 40 bytes in length)
func Repeat(src string) string {
	//bSrc := []byte(src)
	var low, finalKeySize int
	// Determine possible keysizse based on hamming distance, stored as finalKeySize
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
	// Break the ciphertext into blocks of length finalKeySize
	block := make([]string, len(src)/finalKeySize)
	for i := 0; i < len(src)/finalKeySize; i++ {
		block[i] = src[finalKeySize*i : finalKeySize*i+finalKeySize]
	}
	// Transpose those blocks so that there are finalKeySize blocks, all bytes in each index of each block
	// i.e. first block contains the first byte of each block, second block contains second byte of each block, etc...
	block2 := make([]string, finalKeySize)
	for i := 0; i < finalKeySize; i++ {
		var temp = ""
		for j := range block {
			temp += string(block[j][i])
		}
		block2[i] = temp
	}

	// Run single byte xor cracker on the blocks of bytes
	// out := make([]string, len(block2))
	out := ""
	for i := range block2 {
		keyTemp, _ := Xor(block2[i])
		out += keyTemp
	}
	// finalOut := strings.Join(out, "")
	return out
}
