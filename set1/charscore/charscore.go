package charscore

import "strings"

// GetCharScore function to determine "score" of individual characters in potentially decoded message
func GetCharScore(data string) int {
	scores := map[string]int{"e": 27, "t": 26, " ": 25, "a": 24, "o": 23, "i": 22, "n": 21, "s": 20, "h": 19, "r": 18, "d": 17, "l": 16, "c": 15, "u": 14, "m": 13, "w": 12, "f": 11, "g": 10, "y": 9, "p": 8, "b": 7, "v": 6, "k": 5, "j": 4, "x": 3, "q": 2, "z": 1}

	out := scores[data]
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
