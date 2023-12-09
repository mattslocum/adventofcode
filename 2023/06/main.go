package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type race struct {
	time int
	dist int
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	races := parseInput(string(data))
	// fmt.Println(races)

	findWins(races)
}

func parseInput(data string) []race {
	lines := strings.Split(string(data), "\n")
	times := strings.Fields(lines[0])
	dists := strings.Fields(lines[1])

	races := make([]race, len(times)-1)
	for i := 1; i < len(times); i++ {
		t, _ := strconv.Atoi(times[i])
		d, _ := strconv.Atoi(dists[i])
		races[i-1] = race{
			time: t,
			dist: d,
		}
	}

	return races
}

func findWins(races []race) {
	product := 1
	for _, race := range races {
		wins := 0
		for i := 1; i < race.time; i++ {
			dist := (race.time - i) * i
			// println("dist", i, dist)
			if dist > race.dist {
				wins++
			}
		}
		// fmt.Println(wins)
		product *= wins
		// break
	}
	fmt.Println(product)
}
