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
	race := parseInput(string(data))
	fmt.Println(race)

	findWins(race)
}

func parseInput(data string) race {
	lines := strings.Split(string(data), "\n")
	times := strings.Fields(lines[0])
	dists := strings.Fields(lines[1])

	time := ""
	dist := ""
	for i := 1; i < len(times); i++ {
		time += times[i]
		dist += dists[i]
	}

	t, _ := strconv.Atoi(time)
	d, _ := strconv.Atoi(dist)
	return race{
		time: t,
		dist: d,
	}
}

func findWins(race race) {
	wins := 0
	for i := 1; i < race.time; i++ {
		dist := (race.time - i) * i
		// println("dist", i, dist)
		if dist > race.dist {
			wins++
		}
	}
	fmt.Println(wins)
}
