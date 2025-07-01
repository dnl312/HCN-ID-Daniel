package main

import (
	"fmt"
)

var digitToLetters = map[rune][]string{
	'0': {" "},
	'1': {" "},
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func backtrack(index int, digits string, current string, result *[]string) {
	if index == len(digits) {
		*result = append(*result, current)
		return
	}

	letters := digitToLetters[rune(digits[index])]
	for _, letter := range letters {
		backtrack(index+1, digits, current+letter, result)
	}
}

func combinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	for _, d := range digits {
		if d < '0' || d > '9' || d == '1' {
			return []string{}
		}
	}

	var result []string
	backtrack(0, digits, "", &result)
	return result
}

func main() {
	input := "23"
	output := combinations(input)
	fmt.Println("Kombinasi:", output)
}
