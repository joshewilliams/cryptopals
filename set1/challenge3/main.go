package main

import (
	"cryptopals/set1/charscore"
	"cryptopals/set1/xor"
	"encoding/hex"
	"fmt"
)

func main() {
	src := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
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

	fmt.Println(out)
}
