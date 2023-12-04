package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines := strings.Split(string(data), "\n")

	// re := regexp.MustCompile(" +")

	total := 0
	for _, line := range lines {
		scores := strings.Split(line[10:], " | ")
		// winning := re.Split(scores[0], -1)
		// picks := re.Split(scores[1], -1)
		winning := score2Nums(scores[0])
		picks := score2Nums(scores[1])
		sort.Ints(winning)
		sort.Ints(picks)
		total += countMatches(winning, picks)
	}
	fmt.Println(total)
}

func score2Nums(s string) []int {
	nums := []int{}
	vals := strings.Split(s, " ")
	for _, v := range vals {
		num, err := strconv.Atoi(v)
		if err == nil {
			nums = append(nums, num)
		}
	}
	return nums
}

func countMatches(winning, picks []int) int {
	score := 0
	for _, w := range winning {
		if contains[int](picks, w) {
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}
	}
	return score
}

func contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
