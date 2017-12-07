package main

import (
	"bufio"
	"cryptopals/aes"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("./8.txt")
	check(err)
	defer file.Close()

	fscanner := bufio.NewScanner(file)

	for fscanner.Scan() {
		src := fscanner.Text()

	}
}
