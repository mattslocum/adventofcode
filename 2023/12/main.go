package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type springs struct {
	Line string
	Exp  *regexp.Regexp
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	sets := parseInput(string(data))
	// fmt.Println(sets)
	total := 0
	for _, s := range sets {
		count := countPossible(s)
		total += count
	}
	fmt.Println(total)
}

func parseInput(data string) []springs {
	sets := []springs{}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		parts := strings.Split(line, " ")
		broken := strings.Split(parts[1], ",")
		s := springs{Line: parts[0]}
		exp := make([]string, len(broken))
		for pos, val := range broken {
			exp[pos] = fmt.Sprintf("[?#]{%s}", val)
		}
		s.Exp = regexp.MustCompile("^[.]*" + strings.Join(exp, "[.?]+") + "[.]*$")
		sets = append(sets, s)
	}

	return sets
}

func countPossible(s springs) int {
	possible := 0
	// needs more thought. Maybe recursion?
	// for i, r := range s.Line {
	// 	if r == '?' {
	// 		testOn := s.Line[:i] + "." + s.Line[i+1:]
	// 		if s.Exp.MatchString(testOn) {
	// 			fmt.Println(testOn)
	// 			possible++
	// 		}
	// 	}
	// }

	// going with brute force
	wilds := []int{}
	allOn := s.Line
	for i, r := range s.Line {
		if r == '?' {
			wilds = append(wilds, i)
			allOn = allOn[:i] + "." + allOn[i+1:]
		}
	}
	if s.Exp.MatchString(allOn) {
		possible++
	}

	// fmt.Println(wilds)
	all := []string{allOn}
	for _, i := range wilds {
		for _, str := range all {
			off := str[:i] + "#" + str[i+1:]
			all = append(all, off)
			if s.Exp.MatchString(off) {
				possible++
			}
		}
	}
	return possible
}
