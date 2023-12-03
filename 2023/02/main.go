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
	lines := strings.Split(string(data), "\n")

	total := 0
	for _, s := range lines {
		total += checkGame(s)
	}
	fmt.Println(total)
}

func checkGame(line string) int {
	lineSplit := strings.Split(line, ": ")
	game, err := strconv.Atoi(strings.TrimLeft(lineSplit[0], "Game "))
	if err != nil {
		return 0
	}

	gameSets := strings.Split(lineSplit[1], "; ")
	for _, set := range gameSets {
		cubes := strings.Split(set, ", ")
		for _, c := range cubes {
			cube := strings.Split(c, " ")
			num, err := strconv.Atoi(cube[0])
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				os.Exit(1)
			}

			// max 12 red cubes, 13 green cubes, and 14 blue cubes
			switch cube[1] {
			case "red":
				if num > 12 {
					return 0
				}
			case "green":
				if num > 13 {
					return 0
				}
			case "blue":
				if num > 14 {
					return 0
				}
			}
		}
	}

	return game
}
