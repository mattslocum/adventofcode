package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	dest int
	src  int
	size int
}

type elementMap struct {
	src  string
	dest string
	data []Range
}

func (e elementMap) Len() int {
	return len(e.data)
}
func (e elementMap) Less(i, j int) bool {
	return e.data[i].src < e.data[j].src
}
func (e elementMap) Swap(i, j int) {
	e.data[i], e.data[j] = e.data[j], e.data[i]
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines := strings.Split(string(data), "\n")

	seeds := parseSeeds(lines[0])
	fmt.Println(seeds)

	maps := parseMaps(lines[2:])
	findLowest(seeds, maps)
}

func parseSeeds(line string) []int {
	parts := strings.Split(line, " ")
	seeds := []int{}
	for _, val := range parts {
		num, err := strconv.Atoi(val)
		if err == nil {
			seeds = append(seeds, num)
		}
	}
	return seeds
}

func parseMaps(lines []string) map[string]*elementMap {
	m := make(map[string]*elementMap)
	var cur *elementMap
	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.Contains(line, " map:") {
			keys := strings.Split(strings.TrimRight(line, " map:"), "-")
			cur = &elementMap{
				src:  keys[0],
				dest: keys[2],
			}
			m[keys[0]] = cur
			continue
		}

		parts := strings.Split(line, " ")
		dest, _ := strconv.Atoi(parts[0])
		src, _ := strconv.Atoi(parts[1])
		size, _ := strconv.Atoi(parts[2])

		cur.data = append(cur.data, Range{
			dest: dest,
			src:  src,
			size: size,
		})
	}
	for _, el := range m {
		sort.Sort(elementMap(*el))
	}
	return m
}

func findLowest(seeds []int, maps map[string]*elementMap) {
	lowest := math.MaxInt32
	for _, s := range seeds {
		_, val := walkMapLowest("seed", s, maps)
		if val < lowest {
			lowest = val
		}
	}
	fmt.Println(lowest)
}

func walkMapLowest(key string, val int, maps map[string]*elementMap) (string, int) {
	if key == "location" {
		return key, val
	}
	for _, r := range maps[key].data {
		if r.src <= val && r.src+r.size-1 >= val {
			offset := val - r.src
			// fmt.Println("match", key, val, r)
			return walkMapLowest(maps[key].dest, r.dest+offset, maps)
		}
	}
	return walkMapLowest(maps[key].dest, val, maps)
}
