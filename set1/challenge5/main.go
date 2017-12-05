package main

import (
	"cryptopals/set1/xor"
	"fmt"
)

func main() {
	src := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	key := "ICE"
	repeat := true
	hexa := false
	fmt.Println(xor.ToHex(src, key, hexa, repeat))
}
