package charscore

import "strings"
import "math"

// GetCharScore function to determine "score" of individual characters in potentially decoded message
func GetCharScore(data string) int {
	out := 0
	scores := map[string]int{
		"e": 27,
		"3": 27,
		"t": 26,
		"7": 26,
		" ": 25,
		"_": 25,
		"-": 25,
		"a": 24,
		"4": 24,
		"@": 24,
		"o": 23,
		"0": 23,
		"i": 22,
		"1": 22,
		"!": 22,
		"n": 21,
		"s": 20,
		"5": 20,
		"$": 20,
		"h": 19,
		"r": 18,
		"d": 17,
		"l": 16,
		"c": 15,
		"u": 14,
		"m": 13,
		"w": 14,
		"{": 1,
		"}": 1,
		"[": 1,
		"]": 1,
		"(": 1,
		")": 1,
		"f": 11,
		"g": 10,
		"9": 10,
		"6": 10,
		"y": 9,
		"p": 8,
		"b": 7,
		"8": 7,
		"v": 6,
		"k": 5,
		"j": 4,
		"x": 3,
		"q": 2,
		"z": 1,
		"2": 1,
	}

	out = scores[data]
	return out
}

// TotalScore function determines total "score" of given string
func TotalScore(data string) int {
	total := 0
	for _, r := range data {
		total += GetCharScore(strings.ToLower(string(r)))
	}
	return total
}

// Hamming function to calculate Hamming distance between two strings
func Hamming(data1, data2 string) int {
	total := 0
	if data1 == data2 {
		return total
	}
	bData1, bData2 := []byte(data1), []byte(data2)
	for i := 0; i < len(bData1); i++ {
		tmp1, tmp2 := bData1[i], bData2[i]
		for counter := 0; counter < 8; counter++ {
			test := byte(math.Pow(float64(2), float64(counter)))
			if (tmp1 & test) != (tmp2 & test) {
				total++
			}
		}
	}
	return total
}
