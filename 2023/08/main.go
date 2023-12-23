package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines := strings.Split(string(data), "\n")
	steps := strings.Split(lines[0], "")
	lines = lines[2:]

	// Note: according to our data, the Z never happens twice in a loop
	m := parseInput(lines)
	// fmt.Println(m)

	pos := []string{}
	for key, _ := range m {
		if key[2] == 'A' {
			pos = append(pos, key)
		}
	}
	// fmt.Println(pos)

	loopNums := make([]int, len(pos))
	for pi, p := range pos {
		count := 0
		for ; p[2] != 'Z'; count++ {
			i := count % len(steps)
			switch steps[i] {
			case "L":
				p = m[p][0]
			case "R":
				p = m[p][1]
			}
		}
		loopNums[pi] = count
	}
	// fmt.Println(loopNums)
	fmt.Println(lcm(loopNums))
}

func parseInput(lines []string) map[string][]string {
	m := make(map[string][]string)

	for _, line := range lines {
		// hacky parsing that matches the consistent data
		key := line[0:3]
		l := line[7:10]
		r := line[12:15]
		m[key] = []string{l, r}
	}

	return m
}

func lcm(nums []int) int {
	totals := map[int]int{}
	for _, n := range nums {
		factors := primeFactors(n)
		for val, count := range factors {
			if totals[val] < count {
				totals[val] = count
			}
		}
	}

	fac := 1
	for val, count := range totals {
		fac *= int(math.Pow(float64(val), float64(count)))
	}
	return fac
}

func primeFactors(num int) map[int]int {
	factors := map[int]int{}
	for num > 1 {
		if num%2 == 0 {
			factors[2]++
			num = num / 2
			continue
		}
		// TODO: be more efficient on increments of primes
		for n := 3; n <= num; n += 2 {
			if num%n == 0 {
				factors[n]++
				num = num / n
				break
			}
		}
	}
	return factors
}
