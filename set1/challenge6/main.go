package main

import (
	"cryptopals/set1/b64"
	"cryptopals/set1/brute"
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, _ := ioutil.ReadFile("./6.txt")
	s := string(file)
	sAsc := b64.ToASCII(s)
	fmt.Printf(brute.Repeat(sAsc))
}
