package main

import (
	"cryptopals/b64"
	"cryptopals/brute"
	"cryptopals/xor"
	"fmt"
	"io/ioutil"
)

func main() {
	file, _ := ioutil.ReadFile("./6.txt")
	s := string(file)
	sAsc := b64.ToASCII(s)
	key := brute.Repeat(sAsc)
	for i := range key {
		fmt.Println(xor.ToASCII(sAsc, []byte(key[i]), false, true))
		fmt.Printf("\n\n\n")
	}
}
