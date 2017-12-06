package brute

import (
	"cryptopals/set1/charscore"
	"cryptopals/set1/xor"
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
func Repeat(src string) []string {
	keyMap := make(map[int]int)
	keys := make(map[int]int)

	// Determine possible keysizse based on hamming distance, stored as finalKeySize
	for keysize := 2; keysize <= 40; keysize++ {
		tmp1, tmp2, tmp3, tmp4 := src[:keysize], src[keysize:keysize*2], src[keysize*2:keysize*3], src[keysize*3:keysize*4]
		result1 := charscore.Hamming(tmp1, tmp2)
		result2 := charscore.Hamming(tmp2, tmp3)
		result3 := charscore.Hamming(tmp3, tmp4)
		result4 := charscore.Hamming(tmp1, tmp3)
		result5 := charscore.Hamming(tmp1, tmp4)
		result6 := charscore.Hamming(tmp2, tmp4)
		result := (result1 + result2 + result3 + result4 + result5 + result6) / keysize
		keyMap[keysize] = result
	}

	// Grab five keysizes with the smallest hamming distances
	for i := 0; i < 5; i++ {
		var lowKey int
		lowValue := 10000
		for j, k := range keyMap {
			if k < lowValue {
				lowKey = j
				lowValue = k
			}
		}
		keys[lowKey] = lowValue
		delete(keyMap, lowKey)
	}

	// Current issue is that finalKeys (l.82) is out of range at some point during assignment of keyTemp value.
	// My math must be iffy, will straighten this out later.
	var finalKeys []string
	for i := range keys {

		// Break the ciphertext into blocks of length keys[i]
		block := make([]string, len(src)/i)
		for j := 0; j < len(src)/i; j++ {
			block[j] = src[j*i : j*i+i]
		}

		// Transpose those blocks so that there are i blocks, all bytes in each index of each block
		// i.e. first block contains the first byte of each block, second block contains second byte of each block, etc...
		block2 := make([]string, i)
		for j := 0; j < i; j++ {
			var temp = ""
			for k := range block {
				temp += string(block[k][j])
			}
			block2[j] = temp
		}

		// Run single byte xor cracker on the blocks of bytes
		// out := make([]string, len(block2))
		for j, k := range block2 {
			keyTemp, _ := Xor(k)
			finalKeys[j] = keyTemp
		}
	}

	// finalOut := strings.Join(out, "")
	return finalKeys
}
