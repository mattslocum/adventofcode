package main

import (
	"fmt"
	"os"
	"strings"
)

type pattern struct {
	Lines []string
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	patterns := parseInput(string(data))
	total := 0
	for _, pat := range patterns {
		total += findReflection(pat)
	}
	fmt.Println(total)
}

func parseInput(data string) []pattern {
	patterns := []pattern{}
	lines := strings.Split(string(data), "\n")
	pat := pattern{}
	for _, line := range lines {
		if line == "" {
			patterns = append(patterns, pat)
			pat = pattern{}
			continue
		}
		pat.Lines = append(pat.Lines, line)
	}
	patterns = append(patterns, pat)

	return patterns
}

func findReflection(pat pattern) int {
	for i := 1; i < len(pat.Lines); i++ {
		if pat.Lines[i-1] == pat.Lines[i] {
			if isEqual(reverse(pat.Lines[:i]), pat.Lines[i:]) {
				return i * 100
			}
		}
	}
	rot := rotate(pat)
	for i := 1; i < len(rot); i++ {
		if rot[i-1] == rot[i] {
			if isEqual(reverse(rot[:i]), rot[i:]) {
				return i
			}
		}
	}
	return 0
}

func rotate(pat pattern) []string {
	size := len(pat.Lines[0])
	r := make([]string, size)
	for i := 0; i < size; i++ {
		str := ""
		for _, line := range pat.Lines {
			str += string(line[i])
		}
		r[i] = str
	}
	return r
}

func isEqual(a, b []string) bool {
	min := len(a)
	if len(b) < min {
		min = len(b)
	}
	for i := 0; i < min; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func reverse(s []string) []string {
	r := make([]string, len(s))
	copy(r, s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return r
}
