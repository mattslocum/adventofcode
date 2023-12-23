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
	reports := parseInput(string(data))

	sum := 0
	for _, report := range reports {
		sum += findNext(report)
	}
	fmt.Println(sum)

	// part2
	sum = 0
	for _, report := range reports {
		sum += findPrev(report)
	}
	fmt.Println(sum)
}

func parseInput(data string) [][]int {
	lines := strings.Split(string(data), "\n")
	result := make([][]int, len(lines))
	for lineIdx, line := range lines {
		vals := strings.Split(string(line), " ")
		result[lineIdx] = make([]int, len(vals))
		for valIdx, val := range vals {
			i, _ := strconv.Atoi(val)
			result[lineIdx][valIdx] = i
		}
	}

	return result
}

func findNext(line []int) int {
	vals := buildTree(line)
	// fmt.Println(vals)

	for line := len(vals) - 2; line >= 0; line-- {
		last := len(vals[line]) - 1
		lastDeeper := len(vals[line+1]) - 1
		vals[line] = append(vals[line], vals[line][last]+vals[line+1][lastDeeper])
	}
	// fmt.Println("end", vals)
	return vals[0][len(vals[0])-1]
}

func buildTree(line []int) [][]int {
	vals := [][]int{line}
	for depth := 1; ; depth++ {
		prev := vals[depth-1]
		vals = append(vals, make([]int, len(prev)-1))
		isSame := true
		for i := 0; i < len(prev)-1; i++ {
			vals[depth][i] = prev[i+1] - prev[i]
			if i != 0 && isSame {
				isSame = vals[depth][i-1] == vals[depth][i]
			}
		}
		if isSame {
			break
		}
	}
	return vals
}

// Part2
func findPrev(line []int) int {
	vals := buildTree(line)

	for line := len(vals) - 2; line >= 0; line-- {
		vals[line] = append([]int{vals[line][0] - vals[line+1][0]}, vals[line]...)
	}
	return vals[0][0]
}
