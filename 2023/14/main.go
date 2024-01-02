package main

import (
	"fmt"
	"os"
	"slices"
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
	// prep rotated the other direction
	rot := rotate(lines)
	slices.Reverse(rot)

	cache := map[string]int{}
	// large enough loops to find the cache hit
	for i := 0; i < 100000; i++ {
		for j, line := range rot {
			rot[j] = sideO(line)
		}
		slices.Reverse(rot)
		rot = rotate(rot)

		// mod 4 because 4 tilts is 1 cycle
		if (i+1)%4 == 0 {
			cycles := (i + 1) / 4
			key := strings.Join(rot, "")
			if prev, has := cache[key]; has {
				// 100M cycles
				if (1000000000-prev)%(cycles-prev) == 0 {
					count(rot)
					break
				}
			}
			cache[key] = cycles
		}
	}
}

func count(lines []string) {
	size := len(lines[0])
	total := 0
	for _, row := range lines {
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
