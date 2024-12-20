package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(part2())
}

func parseInput(path string) (values []int, gaps []int) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// fmt.Println(string(data))
	// lines := strings.Split(data, "")
	isVal := true
	for _, val := range string(data) {
		i, err := strconv.Atoi(string(val))
		// fmt.Println(val, i)
		if err != nil {
			continue
		}
		if isVal {
			values = append(values, i)
		} else {
			gaps = append(gaps, i)
		}
		isVal = !isVal
	}
	return values, gaps
}

func part2() int {
	values, gaps := parseInput("./09.txt")

	first := values[0]
	values = values[1:]

	defrag := fill(0, first)
	valsUsed := map[int]bool{}

	for i, gap := range gaps {
		for valIdx := len(values) - 1; valIdx >= i; valIdx-- {
			if valsUsed[valIdx] {
				continue
			}
			val := values[valIdx]
			if val <= gap {
				defrag = append(defrag, fill(valIdx+1, val)...)
				valsUsed[valIdx] = true
				gap -= val
			}
			if gap == 0 {
				break
			}
		}
		if gap > 0 {
			defrag = append(defrag, fill(0, gap)...)
		}
		if len(values) > i {
			if !valsUsed[i] {
				// +1 because we already removed the first value before starting
				defrag = append(defrag, fill(i+1, values[i])...)
			} else {
				// moved val now creates a gap that can't be filled
				defrag = append(defrag, fill(0, values[i])...)
			}
		}
	}
	// fmt.Println(defrag)

	answer := 0
	for i, val := range defrag {
		answer += i * val
	}

	return answer
}

func fill(val int, count int) []int {
	var ret []int
	for i := 0; i < count; i++ {
		ret = append(ret, val)
	}
	return ret
}
