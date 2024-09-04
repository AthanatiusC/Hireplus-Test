package main

import (
	"fmt"
	"strings"
)

func main() {
	var limit int
	fmt.Println("Masukan Limit(n) : ")
	fmt.Scanln(&limit)

	wordMap := make(map[string][]int)
	for i := 1; i <= limit; i++ {
		var word string
		fmt.Scanln(&word)

		wordMap[strings.ToLower(word)] = append(wordMap[strings.ToLower(word)], i)
	}

	var result []int
	for _, val := range wordMap {
		if len(val) > 1 {
			result = append(result, val...)
		}
	}

	if len(result) != 0 {
		fmt.Println(strings.Trim(strings.Replace(fmt.Sprint(result), " ", " ", -1), "[]"))
	} else {
		fmt.Println("false")
	}
}
