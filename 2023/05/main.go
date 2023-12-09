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
	offset int
	src    int
	end    int
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
	// fmt.Println(seeds)

	// parseMaps(lines[2:])
	maps := parseMaps(lines[2:])
	findLowest(seeds, maps)
}

func parseSeeds(line string) [][]int {
	parts := strings.Split(line, " ")
	seeds := [][]int{}
	start := 0 // we never start at 0
	for _, val := range parts {
		num, err := strconv.Atoi(val)
		if err == nil {
			if start == 0 {
				start = num
			} else {
				seeds = append(seeds, []int{start, start + num - 1})
				start = 0
			}
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
			offset: dest - src,
			src:    src,
			end:    src + size - 1,
		})
	}
	for _, el := range m {
		sort.Sort(elementMap(*el))
		// do gaps even happen?
		gaps := []Range{}
		for i, row := range el.data {
			if i < len(el.data)-1 {
				next := el.data[i+1]
				if row.end+1 < next.src {
					gaps = append(gaps, Range{
						offset: 0,
						src:    row.end + 1,
						end:    next.src - 1,
					})
				}
			}
		}
		if len(gaps) > 0 {
			el.data = append(el.data, gaps...)
			sort.Sort(elementMap(*el))
		}
		if el.data[0].src != 0 {
			el.data = append([]Range{{
				offset: 0,
				src:    0,
				end:    el.data[0].src - 1,
			}}, el.data...)
		}
	}
	return m
}

type itemRange struct {
	start int
	end   int
	level string
}

func findLowest(seeds [][]int, maps map[string]*elementMap) {
	items := make([]itemRange, len(seeds))
	for i, s := range seeds {
		items[i] = itemRange{
			start: s[0],
			end:   s[1],
			level: "seed",
		}
	}

	locations := []int{}
	for i := 0; i < len(items); i++ {
		item := items[i]
		// fmt.Println(item)
		if item.level == "location" {
			// we found the end
			locations = append(locations, item.start)
			continue
		}

		for rangeI, span := range maps[item.level].data {
			if span.src <= item.start && span.end >= item.start {
				end := min(item.end, span.end)
				items = append(items, itemRange{
					start: item.start + span.offset,
					end:   end + span.offset,
					level: maps[item.level].dest,
				})
				if end < item.end {
					// item not finished, set the new start and keep looping
					item.start = end + 1
				} else {
					break // Full match on this item
				}
			}
			if rangeI == len(maps[item.level].data)-1 {
				// We are past the ranges, map straight through
				items = append(items, itemRange{
					start: item.start,
					end:   item.end,
					level: maps[item.level].dest,
				})
			}
		}
	}
	// fmt.Println(locations)

	lowest := math.MaxInt32
	for _, num := range locations {
		if num < lowest {
			lowest = num
		}
	}
	fmt.Println(lowest)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
