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

	elves := []int{}
	elfCal := 0
	for _, s := range lines {
		if s == "" {
			elves = append(elves, elfCal)
			elfCal = 0
		} else {
			num, _ := strconv.Atoi(s)
			elfCal += num
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elves)))
	fmt.Println(elves)
	top := 0
	for i := 0; i < 3; i++ {
		top += elves[i]
	}
	fmt.Println(top)
}
