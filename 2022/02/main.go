package main

import (
	"fmt"
	"os"
	"strings"
)

type plays int

const (
	Rock plays = iota
	Paper
	Scissors
)

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines := strings.Split(string(data), "\n")

	plays := map[string]plays{
		"A": Rock,
		"B": Paper,
		"C": Scissors,
		"X": Rock,
		"Y": Paper,
		"Z": Scissors,
	}
	scores := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	score := 0
	for _, s := range lines {
		votes := strings.Split(string(s), " ")

		score += scores[votes[1]]
		score += shoot(plays[votes[0]], plays[votes[1]])
	}
	fmt.Println(score)
}

func shoot(op, me plays) int {
	// Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock
	if me == Rock {
		switch op {
		case Scissors:
			return 6
		case Paper:
			return 0
		}
	}
	if me == Paper {
		switch op {
		case Rock:
			return 6
		case Scissors:
			return 0
		}
	}
	if me == Scissors {
		switch op {
		case Paper:
			return 6
		case Rock:
			return 0
		}
	}
	// tie
	return 3
}
