package main

import (
	"cryptopals/aes"
	"cryptopals/b64"
	"fmt"
	"io/ioutil"
)

func main() {
	file, _ := ioutil.ReadFile("./7.txt")
	s := string(file)
	bAsc := []byte(b64.ToASCII(s))
	key := []byte("YELLOW SUBMARINE")

	fmt.Println(string(aes.ECB(bAsc, key)))
}
