package main

import (
	"cryptopals/set1/b64"
	//"encoding/hex"
	"fmt"
)

func main() {
	string1 := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	fmt.Println(b64.Decode(string1))
}
