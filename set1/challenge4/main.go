package main

import (
	"bufio"
	"cryptopals/set1/brute"
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

	fscanner := bufio.NewScanner(file)
	var highest int
	var finalResult string
	for fscanner.Scan() {
		src := fscanner.Text()
		result, high := brute.Xor(src)
		if high > highest {
			highest = high
			finalResult = result
		}
	}
	fmt.Print(finalResult)
}
