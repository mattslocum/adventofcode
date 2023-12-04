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

	cards := []int{}
	for _, line := range lines {
		scores := strings.Split(line[10:], " | ")
		winning := score2Nums(scores[0])
		picks := score2Nums(scores[1])
		// TODO: If we cared about performance, sort and then short circuit the search later so this actually helps
		sort.Ints(winning)
		sort.Ints(picks)
		cards = append(cards, countMatches(winning, picks))
	}

	// Pre-sum each card from end to start to know it's value
	// that way when we say it gets used by a prior card, we already know the value.
	points := make([]int, len(cards))
	for i := len(cards) - 1; i >= 0; i-- {
		sweeping := 1
		for x := 1; x <= cards[i]; x++ {
			sweeping += points[i+x]
		}
		points[i] = sweeping
	}
	// fmt.Println(points)

	total := 0
	for _, n := range points {
		total += n
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
			score++
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
