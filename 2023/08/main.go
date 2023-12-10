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
	lines := strings.Split(string(data), "\n")
	steps := strings.Split(lines[0], "")
	lines = lines[2:]

	m := parseInput(lines)
	// fmt.Println(m)

	count := 0
	pos := "AAA"
	for ; ; count++ {
		i := count % len(steps)
		switch steps[i] {
		case "L":
			pos = m[pos][0]
		case "R":
			pos = m[pos][1]
		}
		if pos == "ZZZ" {
			count++
			break
		}
	}
	fmt.Println(count)
}

func parseInput(lines []string) map[string][]string {
	m := make(map[string][]string)

	for _, line := range lines {
		// hacky parsing that matches the consistent data
		key := line[0:3]
		l := line[7:10]
		r := line[12:15]
		m[key] = []string{l, r}
	}

	return m
}
