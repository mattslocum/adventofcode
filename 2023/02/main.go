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
	_, err := strconv.Atoi(strings.TrimLeft(lineSplit[0], "Game "))
	if err != nil {
		return 0
	}

	var (
		minRed   = 0
		minGreen = 0
		minBlue  = 0
	)

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

			switch cube[1] {
			case "red":
				if num > minRed {
					minRed = num
				}
			case "green":
				if num > minGreen {
					minGreen = num
				}
			case "blue":
				if num > minBlue {
					minBlue = num
				}
			}
		}
	}

	return minRed * minGreen * minBlue
}
