package main

import (
	"bufio"
	"cryptopals/brute"
	"cryptopals/xor"
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
	defer file.Close()

	fscanner := bufio.NewScanner(file)
	var highest int
	var finalResult, finalOutput, ciphertext string

	for fscanner.Scan() {
		src := fscanner.Text()
		result, high := brute.Xor(src, true)
		if high > highest {
			highest = high
			finalResult = result
			ciphertext = src
			finalOutput = xor.ToASCII(src, finalResult, true, false)
		}
	}
	fmt.Println("Key:", finalResult)
	fmt.Println("Ciphertext:", ciphertext)
	fmt.Println("Plaintext:", finalOutput)
}
