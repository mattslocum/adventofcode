package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// replacer is step 2
	replacer := strings.NewReplacer(
		"zero", "0",
		"one", "1",
		"two", "2",
		"three", "3",
		"four", "4",
		"five", "5",
		"six", "6",
		"seven", "7",
		"eight", "8",
		"nine", "9",
	)

	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	vals := strings.Split(string(data), "\n")

	total := 0
	for _, s := range vals {
		total += findCode(replacer.Replace(s))
	}
	fmt.Println(total)
}

func findCode(s string) int {
	var first, second string
	for _, c := range s {
		if _, err := strconv.Atoi(string(c)); err == nil {
			first = string(c)
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if _, err := strconv.Atoi(string(s[i])); err == nil {
			second = string(s[i])
			break
		}
	}
	if i, err := strconv.Atoi(first + second); err == nil {
		return i
	}
	return 0
}
