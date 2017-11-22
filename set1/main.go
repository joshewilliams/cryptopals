package main

import (
	"cryptopals/set1/xor"
	"fmt"
)

func main() {
	string1 := "SSdtIGtpbGxpbmcgeW91ci"
	out := b64.Decode(string1)
	fmt.Println(out)
}
