package main

import (
	"bufio"
	"cryptopals/brute"
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
