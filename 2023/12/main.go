package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type springs struct {
	Line   string
	Sets   []string
	Broken []int
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	sets := parseInput(string(data))
	total := 0
	for _, s := range sets {
		// fmt.Println(s.Sets, s.Broken)
		count := countPossible(s, &map[string]int{})
		// fmt.Println(count)
		total += count
	}
	fmt.Println(total)
}

func parseInput(data string) []springs {
	sets := []springs{}
	lines := strings.Split(string(data), "\n")
	re := regexp.MustCompile("[.]+")
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		parts[0] = strings.Repeat(parts[0]+"?", 5)
		parts[0] = parts[0][:len(parts[0])-1] // trim ending ?
		parts[1] = strings.Repeat(parts[1]+",", 5)
		parts[1] = parts[1][:len(parts[1])-1] // trim ending ,
		broken := strings.Split(parts[1], ",")
		s := springs{
			Line:   parts[0],
			Broken: make([]int, len(broken)),
			Sets:   re.Split(strings.Trim(parts[0], "."), -1),
		}
		for pos, val := range broken {
			i, _ := strconv.Atoi(val)
			s.Broken[pos] = i
		}
		sets = append(sets, s)
	}

	return sets
}

func countPossible(s springs, cache *map[string]int) int {
	key := fmt.Sprintf("%s,%d", strings.Join(s.Sets, ""), sum(s.Broken...))
	if val, has := (*cache)[key]; has {
		return val
	}

	possible := 0
	if len(s.Broken) == 0 {
		// if we are out of broken matches make sure that no more are required
		for _, set := range s.Sets {
			if strings.Contains(set, "#") {
				return 0
			}
		}
		return 1
	}
	if len(s.Sets) == 0 {
		// if we are out of sets, make sure we don't need any more
		if len(s.Broken) == 0 {
			return 1 // good match
		}
		return 0
	}

	l := s.Broken[0]
	sumBroken := sum(s.Broken[1:]...)
	lenFirst := len(s.Sets[0])
	for i := 0; i+l <= lenFirst; i++ {
		if i > 0 && s.Sets[0][i-1] == '#' {
			break // previous can't be # on itteration
		}
		set := make([]string, len(s.Sets))
		copy(set, s.Sets)
		set[0] = set[0][i+l:]
		// at end or optional as next
		if set[0] == "" || set[0][0] == '?' {
			if set[0] == "" {
				set = set[1:] // remove empty set
			} else if set[0][0] == '?' {
				set[0] = set[0][1:] // remove extra for the space
				if set[0] == "" {
					set = set[1:] // remove empty set
				}
			}
			// optimize by stopping impossible loops early
			if sumBroken > lenStrs(set...) {
				break
			}
			possible += countPossible(springs{
				Line:   s.Line,
				Broken: s.Broken[1:],
				Sets:   set,
			}, cache)
		}
	}
	// if broken isn't required, assume we matched none
	if !strings.Contains(s.Sets[0], "#") {
		possible += countPossible(springs{
			Line:   s.Line,
			Broken: s.Broken,
			Sets:   s.Sets[1:],
		}, cache)
	}

	(*cache)[key] = possible
	return possible
}

func sum(vals ...int) int {
	s := 0
	for _, v := range vals {
		s += v
	}
	return s
}

func lenStrs(strs ...string) int {
	s := 0
	for _, v := range strs {
		s += len(v)
	}
	return s
}
