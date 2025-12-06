package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(part2())
}

var cache11 = map[int][]int{}

func parseInput(path string) []int {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	data = []byte(strings.Trim(string(data), "\n"))
	nums := strings.Split(string(data), " ")
	var values []int
	for _, val := range nums {
		i, err := strconv.Atoi(string(val))
		if err == nil {
			values = append(values, i)
		}
	}
	return values
}

func part2() int {
	nums := parseInput("./11Test.txt")
	fmt.Println("nums", nums)

	for i := 0; i < 75; i++ {
		fmt.Println(i, len(nums))
		nums = runMutation(nums)
	}
	// fmt.Println(nums)

	return len(nums)
}

/*
*
Initial arrangement:
125 17

After 1 blink:
253000 1 7

After 2 blinks:
253 0 2024 14168

After 3 blinks:
512072 1 20 24 28676032

After 4 blinks:
512 72 2024 2 0 2 4 2867 6032

After 5 blinks:
1036288 7 2 20 24 4048 1 4048 8096 28 67 60 32

After 6 blinks:
2097446912 14168 4048 2 0 2 4 40 48 2024 40 48 80 96 2 8 6 7 6 0 3 2
*/
func runMutation(nums []int) []int {
	var list []int

	for _, val := range nums {
		// fmt.Println(val, cache11)
		if cache11[val] == nil {
			digits := len(fmt.Sprintf("%d", val))
			var result []int
			if val == 0 {
				result = append(result, 1)
			} else if digits%2 == 0 {
				// let digitSplitter = pow(base: 10, pow: digits / 2)
				digitSplitter := pow(10, digits/2)
				// first half of the digits.
				result = append(result, val/digitSplitter)
				result = append(result, val%digitSplitter)
			} else {
				result = append(result, val*2024)
			}
			cache11[val] = result
		}
		list = append(list, cache11[val]...)
	}

	return list
}

// TODO: replace with native
func pow(base int, pow int) int {
	var answer = 1
	for range pow {
		answer *= base
	}
	return answer
}
