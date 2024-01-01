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
	lines := parseInput(string(data))
	sumNorth(lines)
}

func parseInput(data string) []string {
	return strings.Split(string(data), "\n")
}

func sumNorth(lines []string) {
	rot := rotate(lines)
	// slide left, which is north
	for i, line := range rot {
		rot[i] = sideO(line)
	}

	size := len(rot[0])
	total := 0
	for _, row := range rot {
		for i, c := range row {
			if c == 'O' {
				total += size - i
			}
		}
	}
	fmt.Println(total)
}

func rotate(lines []string) []string {
	size := len(lines[0])
	r := make([]string, size)
	for i := 0; i < size; i++ {
		str := ""
		for _, line := range lines {
			str += string(line[i])
		}
		r[i] = str
	}
	return r
}

func sideO(line string) string {
	str := []byte(line)
	for i := 0; i < len(line); i++ {
		if line[i] == 'O' {
			for back := i - 1; back >= 0; back-- {
				if str[back] != '.' || (back == 0 && str[back] == '.') {
					if back == 0 && str[back] == '.' {
						back = -1
					}
					str[i] = '.'
					str[back+1] = 'O'
					break
				}
			}
		}
	}
	return string(str)
}
