package main

import (
	"cryptopals/brute"
	"cryptopals/xor"
	"encoding/hex"
	"fmt"
)

func main() {
	s := "d3d9d1c2ddd6e3efaceaf5ede8c7acc7f9c7e0a8eac7a9e5"
	bAsc, _ := hex.DecodeString(s)
	sAsc := string(bAsc[:])
	key := brute.Repeat(sAsc)
	for i := range key {
		fmt.Printf("%x\n", key[i])
		fmt.Println(xor.ToASCII(sAsc, []byte(key[i]), false, true))
		fmt.Printf("\n\n")
	}
}
