package main

import (
	"fmt"
)

func isValid(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{
		'<': '>',
		'{': '}',
		'[': ']',
	}

	for _, char := range s {
		if _, ok := pairs[char]; ok {
			stack = append(stack, char)
		} else if len(stack) > 0 && char == pairs[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}

func main() {
	var input string
	fmt.Println("Masukan Input : ")
	fmt.Scanln(&input)

	if isValid(input) {
		fmt.Println("TRUE")
	} else {
		fmt.Println("FALSE")
	}
}
