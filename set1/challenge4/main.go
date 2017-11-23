package main

import (
	"bufio"
	"cryptopals/set1/charscore"
	"cryptopals/set1/xor"
	"encoding/hex"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("./4.txt")
	check(err)

	fscanner := bufio.NewScanner(file)
	hightop := 0
	var resulttop string
	for fscanner.Scan() {
		src := fscanner.Text()
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
		if high > hightop {
			hightop = high
			resulttop = out
		}

	}

	fmt.Print(resulttop)
}
