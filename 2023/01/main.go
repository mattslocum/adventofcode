package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	vals := strings.Split(string(data), "\n")

	total := 0
	for _, s := range vals {
		total += findCode(s)
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
