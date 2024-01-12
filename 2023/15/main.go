package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	values := parseInput(string(data))
	total := 0
	for _, str := range values {
		h := hash(str)
		total += h
	}
	fmt.Println(total)
}

func parseInput(data string) []string {
	strings := strings.Split(string(data), ",")
	return strings
}

func hash(str string) int {
	h := 0
	for _, c := range str {
		h += int(c)
		h *= 17
		h %= 256
	}
	return h
}
