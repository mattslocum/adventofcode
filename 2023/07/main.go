package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
	cards []int
	bid   int
	rank  rank
}

type games struct {
	hands []hand
}

func (h games) Len() int {
	return len(h.hands)
}
func (h games) Less(i, j int) bool {
	if h.hands[i].rank != h.hands[j].rank {
		return h.hands[i].rank < h.hands[j].rank
	}
	for x, _ := range h.hands[i].cards {
		iCard := h.hands[i].cards
		jCard := h.hands[j].cards
		if iCard[x] != jCard[x] {
			return h.hands[i].cards[x] < h.hands[j].cards[x]
		}
	}
	panic("this should never happen")
}
func (e games) Swap(i, j int) {
	e.hands[i], e.hands[j] = e.hands[j], e.hands[i]
}

var scores = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

type rank int

const (
	card rank = iota
	kind2
	pair2
	kind3
	fullHouse
	kind4
	kind5
)

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	game := games{
		hands: parseInput(string(data)),
	}
	// sorts lowest to highest score
	sort.Sort(game)
	// fmt.Println(game)
	score := 0
	for i := len(game.hands) - 1; i > -1; i-- {
		score += game.hands[i].bid * (i + 1)
	}
	fmt.Println(score)
}

func parseInput(data string) []hand {
	lines := strings.Split(string(data), "\n")
	hands := make([]hand, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, " ")
		bid, _ := strconv.Atoi(parts[1])
		cards := strings.Split(parts[0], "")
		cardsI := make([]int, 5)
		for j, c := range cards {
			cardsI[j] = scores[c]
		}
		hands[i] = hand{
			cards: cardsI,
			bid:   bid,
			rank:  findRank(cards),
		}
	}

	return hands
}

func findRank(cards []string) rank {
	matches := map[string]int{}
	for _, c := range cards {
		matches[c]++
	}
	switch len(matches) {
	case 1:
		return kind5
	case 2: // either kind4 or fullHouse
		for _, num := range matches {
			if num == 4 || num == 1 {
				return kind4
			}
			return fullHouse
		}
	case 3: // either kind3 or pair2
		for _, num := range matches {
			if num == 3 {
				return kind3
			}
		}
		return pair2
	case 4:
		return kind2
	}
	return card
}
