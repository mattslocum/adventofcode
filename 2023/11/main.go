package main

import (
	"fmt"
	"os"
	"strings"
)

type galaxy struct {
	X int
	Y int
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	galaxies := parseInput(string(data))
	// fmt.Println(galaxies)
	sumDistances(galaxies)
}

func parseInput(data string) []galaxy {
	galaxies := []galaxy{}
	lines := strings.Split(string(data), "\n")
	// expand X gaps
	for x := 0; x < len(lines[0]); x++ {
		for y, line := range lines {
			if line[x] == '#' {
				break
			}
			// reached end?
			if y == len(lines)-1 {
				for y2, _ := range lines {
					lines[y2] = lines[y2][:x+1] + lines[y2][x:]
				}
				x++ // because we added an extra column
			}
		}
	}
	// find galaxies
	yAdd := 0
	for y, line := range lines {
		foundGxy := false
		for x, val := range line {
			if val == '#' {
				foundGxy = true
				galaxies = append(galaxies, galaxy{X: x, Y: y + yAdd})
			}
		}
		if !foundGxy {
			yAdd++
		}
	}

	return galaxies
}

func sumDistances(galaxies []galaxy) {
	sum := 0
	for i, gal := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			sum += abs(gal.X - galaxies[j].X)
			sum += abs(gal.Y - galaxies[j].Y)
		}
	}
	fmt.Println(sum)
}

func abs(val int) int {
	if val < 0 {
		return val * -1
	}
	return val
}
