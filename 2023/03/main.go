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
		for x := 0; x < len(line); x++ {
			if lines[y][x] == '*' {
				total += findGeer(lines, x, y)
			}
		}
	}
	fmt.Println(total)
}

func findGeer(lines []string, xPos int, yPos int) int {
	maxY := len(lines)
	maxX := len(lines[0]) // assuming perfect square
	nums := []int{}
	for y := max(0, yPos-1); y <= yPos+1 && y < maxY; y++ {
		prevWasNum := false
		for x := max(0, xPos-1); x <= xPos+1 && x < maxX; x++ {
			if y == yPos && x == xPos {
				prevWasNum = false
				// Skip the center
				continue
			}
			if unicode.IsDigit(rune(lines[y][x])) {
				if !prevWasNum {
					nums = append(nums, findNum(lines[y], x))
					prevWasNum = true
				}
			} else {
				prevWasNum = false
			}
			if len(nums) == 2 {
				// fmt.Println(nums)
				return nums[0] * nums[1]
			}
		}
	}
	return 0
}

func findNum(s string, x int) int {
	digits := []byte{s[x]}
	for d := x - 1; d >= 0; d-- {
		if unicode.IsDigit(rune(s[d])) {
			digits = append([]byte{s[d]}, digits...)
		} else {
			break
		}
	}
	for d := x + 1; d < len(s); d++ {
		if unicode.IsDigit(rune(s[d])) {
			digits = append(digits, s[d])
		} else {
			break
		}
	}
	num, _ := strconv.Atoi(string(digits))
	return num
}
