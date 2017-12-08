package main

import (
	"bufio"
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
	total := 0
	out := ""

	fscanner := bufio.NewScanner(file)

	for fscanner.Scan() {
		src := fscanner.Text()
		var blocks []string
		tmp := 0

		for i := 0; i < len(src)/32; i++ {
			blocks = append(blocks, src[i*32:(i+1)*32])
		}

		for i, j := range blocks {
			for k, l := range blocks {
				if i != k {
					if j == l {
						tmp++
					}
				}
			}
		}
		if tmp > total {
			total = tmp
			out = src
		}
	}

	fmt.Println(out)
}
