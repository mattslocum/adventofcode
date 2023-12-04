package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines := strings.Split(string(data), "\n")

	total := 0
	for y, line := range lines {
		num := ""
		for x := 0; x < len(line); x++ {
			if unicode.IsDigit(rune(lines[y][x])) {
				num += string(lines[y][x])
			} else {
				if num != "" {
					// We have a number that just ended
					total += checkPart(lines, num, x-len(num), y)
					// if num == "123" {
					// 	os.Exit(0)
					// }
					num = ""
				}
			}
		}
		if num != "" {
			// We have a number at the end of the line
			total += checkPart(lines, num, len(lines[y])-len(num), y)
			num = ""
		}
	}
	fmt.Println(total)
}

func checkPart(lines []string, num string, xPos int, yPos int) int {
	maxY := len(lines)
	maxX := len(lines[0])
	// start top left
	for y := max(0, yPos-1); y <= yPos+1 && y < maxY; y++ {
		for x := max(0, xPos-1); x <= xPos+len(num) && x < maxX; x++ {
			if y == yPos && x == xPos {
				// Skip the number
				x += len(num)
				if x >= maxX {
					// Don't go over the edge
					continue
				}
			}
			// fmt.Println(string(lines[y][x]), y, x)
			if lines[y][x] != '.' {
				// fmt.Println("found " + num)
				part, _ := strconv.Atoi(num)
				return part
			}
		}
	}
	return 0
}
