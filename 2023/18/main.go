package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type instruction struct {
	Dir   string
	Steps int
}

type grid struct {
	Hor  segments
	Vert segments
}

func main() {
	data, err := os.ReadFile("./test.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	insts := parseInput(string(data))
	g := walk(insts)

	// split vertical segments by horizontal lines
	for _, line := range g.Hor.Segs {
		g.Vert.Split(line.Pos)
	}
	// split horizontal segments by vertical lines
	for _, line := range g.Vert.Segs {
		g.Hor.Split(line.Pos)
	}

	calcArea(g)
}

func parseInput(data string) []instruction {
	lines := strings.Split(string(data), "\n")
	instructions := make([]instruction, len(lines))
	for i, line := range lines {
		parts := strings.Split(string(line), " ")
		steps, _ := strconv.Atoi(parts[1])
		instructions[i] = instruction{
			Dir:   parts[0],
			Steps: steps,
		}
	}
	return instructions
}

func walk(insts []instruction) grid {
	var (
		curY, curX, maxY, maxX, minY, minX int
	)
	g := grid{}
	for _, inst := range insts {
		switch inst.Dir {
		case "R":
			g.Hor.Add(segment{
				Pos:   curY,
				Start: curX,
				End:   curX + inst.Steps,
			})
			curX += inst.Steps
			if curX > maxX {
				maxX = curX
			}
		case "L":
			g.Hor.Add(segment{
				Pos:   curY,
				Start: curX - inst.Steps,
				End:   curX,
			})
			curX -= inst.Steps
			if curX < minX {
				minX = curX
			}
		case "U":
			g.Vert.Add(segment{
				Pos:   curX,
				Start: curY - inst.Steps,
				End:   curY,
			})
			curY -= inst.Steps
			if curY < minY {
				minY = curY
			}
		case "D":
			g.Vert.Add(segment{
				Pos:   curX,
				Start: curY,
				End:   curY + inst.Steps,
			})
			curY += inst.Steps
			if curY > maxY {
				maxY = curY
			}
		}
	}

	return g
}

func calcArea(g grid) {
	total := 0
	vert := g.Vert.UniqueStarts()
	hor := g.Hor.UniqueStarts()
	for _, y := range vert {
		// FIX: only go 1 layer
		fmt.Println("row", y.Start, y.End-y.Start-1)
		for xi, x := range hor {
			// top side
			// total += x.End - x.Start
			if !g.Vert.HasOddLeft(x.Start, y.Start) {
				fmt.Println("skipping", y.Start, x.Start)
				// Add skipped X value
				continue
			}
			// left side, but exclude top since we already counted it
			// total += y.End - y.Start - 1
			area := (y.End - y.Start - 1) * (x.End - x.Start - 1)
			fmt.Println(area, x.End-x.Start)
			total += area
			if xi+1 == len(vert) {
				fmt.Println("add edge")
				// top side
				// total += x.End - x.Start
				// left side, but exclude top since we already counted it
				// total += y.End - y.Start - 1
			}
		}
		// total += y.End - y.Start - 1
	}
	fmt.Println("total", total)
}

func writeData(field map[int]map[int]string, maxY, maxX, minY, minX int) {
	out := ""
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if val, has := field[y][x]; has {
				out += val
			} else {
				out += "."
			}
		}
		out += "\n"
	}

	os.WriteFile("./result.txt", []byte(out), 0644)
}

type segment struct {
	Pos   int // Y for Horisontal, and X for Vertical
	Start int
	End   int
}

type segments struct {
	Segs []segment
}

func (s *segments) Add(seg segment) {
	s.Segs = append(s.Segs, seg)
	sort.Sort(s)
}

func (s *segments) Split(pos int) {
	add := []segment{}
	for i, seg := range s.Segs {
		if seg.Start < pos && seg.End > pos {
			add = append(add, segment{Pos: seg.Pos, Start: pos, End: seg.End})
			s.Segs[i].End = pos - 1
		}
	}
	// TODO: be more efficent
	if len(add) > 0 {
		s.Segs = append(s.Segs, add...)
		sort.Sort(s)
	}
}

func (s *segments) Len() int      { return len(s.Segs) }
func (s *segments) Swap(i, j int) { s.Segs[i], s.Segs[j] = s.Segs[j], s.Segs[i] }
func (s *segments) Less(i, j int) bool {
	if s.Segs[i].Start == s.Segs[j].Start {
		return s.Segs[i].Pos < s.Segs[j].Pos
	}
	return s.Segs[i].Start < s.Segs[j].Start
}

// Returns a unique set of starts. Pos results will be the first found in each start
func (s *segments) UniqueStarts() []segment {
	result := []segment{}
	found := map[int]bool{}
	for _, seg := range s.Segs {
		if _, has := found[seg.Start]; !has {
			found[seg.Start] = true
			result = append(result, seg)
		}
	}
	return result
}

func (s *segments) HasOddLeft(pos, depth int) bool {
	left := 0
	for _, seg := range s.Segs {
		if pos >= seg.Pos && depth >= seg.Start && depth < seg.End {
			left++
		}
		// TODO: break loop after we are past what is possible
	}
	return left%2 == 1
}
