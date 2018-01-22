package main

import (
	"cryptopals/brute"
	"cryptopals/xor"
	"encoding/hex"
	"fmt"
)

func main() {
	s := "39090249051f11084f071e5b4722223a2a2a2115035110361440025c300d5d510c580d142f185e1a053e1f59110e"
	bAsc, _ := hex.DecodeString(s)
	sAsc := string(bAsc[:])
	key := brute.Repeat(sAsc)
	for i := range key {
		fmt.Printf("%x\n", key[i])
		fmt.Println(xor.ToASCII(sAsc, []byte(key[i]), false, true))
		fmt.Printf("\n\n")
	}
}
