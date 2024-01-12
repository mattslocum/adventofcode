package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type dir int

const (
	FromL dir = iota
	FromR
	FromT
	FromB
)

type cell struct {
	Val    int
	Visits map[dir]map[int]int // direction map of dir count that maps to step count
	Steps  int                 // just for printing fun
}

type step struct {
	Count int // distance in current direction
	From  dir
	Steps int
	X     int
	Y     int
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cells := parseInput(string(data))
	walk(cells)
}

func parseInput(data string) [][]cell {
	lines := strings.Split(string(data), "\n")
	cells := make([][]cell, len(lines))
	for y, line := range lines {
		chars := strings.Split(line, "")
		cells[y] = make([]cell, len(chars))
		for x, c := range chars {
			i, _ := strconv.Atoi(c)
			cells[y][x] = cell{Val: i, Visits: map[dir]map[int]int{}}
		}
	}

	return cells
}

func walk(data [][]cell) {
	todo := []step{
		{Y: 0, X: 1, From: FromL, Count: 1, Steps: data[0][1].Val},
		{Y: 1, X: 0, From: FromT, Count: 1, Steps: data[1][0].Val},
	}
	total := 0
	minEnd := 99999999
	for len(todo) > 0 {
		total++
		st := todo[0]
		todo = todo[1:]
		cur := &data[st.Y][st.X]

		if st.Y == len(data)-1 && st.X == len(data[0])-1 {
			if st.Steps < minEnd {
				minEnd = st.Steps
			}
			continue
		}

		// process current item
		if cur.Visits[st.From][st.Count] == 0 || st.Steps < cur.Visits[st.From][st.Count] {
			if _, ok := cur.Visits[st.From]; !ok {
				cur.Visits[st.From] = map[int]int{}
			}
			cur.Visits[st.From][st.Count] = st.Steps
			if cur.Steps == 0 || st.Steps < cur.Steps {
				cur.Steps = st.Steps
			}
		} else {
			continue
		}

		// breadth first search
		switch st.From {
		case FromL:
			// straight
			if st.Count < 3 && st.X < len(data[0])-1 {
				todo = append(todo, step{Y: st.Y, X: st.X + 1, From: FromL, Count: st.Count + 1, Steps: st.Steps + data[st.Y][st.X+1].Val})
			}
			// up
			if st.Y > 0 {
				todo = append(todo, step{Y: st.Y - 1, X: st.X, From: FromB, Count: 1, Steps: st.Steps + data[st.Y-1][st.X].Val})
			}
			// down
			if st.Y < len(data)-1 {
				todo = append(todo, step{Y: st.Y + 1, X: st.X, From: FromT, Count: 1, Steps: st.Steps + data[st.Y+1][st.X].Val})
			}
		case FromR:
			// straight
			if st.Count < 3 && st.X > 0 {
				todo = append(todo, step{Y: st.Y, X: st.X - 1, From: FromR, Count: st.Count + 1, Steps: st.Steps + data[st.Y][st.X-1].Val})
			}
			// up
			if st.Y > 0 {
				todo = append(todo, step{Y: st.Y - 1, X: st.X, From: FromB, Count: 1, Steps: st.Steps + data[st.Y-1][st.X].Val})
			}
			// down
			if st.Y < len(data)-1 {
				todo = append(todo, step{Y: st.Y + 1, X: st.X, From: FromT, Count: 1, Steps: st.Steps + data[st.Y+1][st.X].Val})
			}
		case FromT:
			// straight
			if st.Count < 3 && st.Y < len(data)-1 {
				todo = append(todo, step{Y: st.Y + 1, X: st.X, From: FromT, Count: st.Count + 1, Steps: st.Steps + data[st.Y+1][st.X].Val})
			}
			// left
			if st.X > 0 {
				todo = append(todo, step{Y: st.Y, X: st.X - 1, From: FromR, Count: 1, Steps: st.Steps + data[st.Y][st.X-1].Val})
			}
			// right
			if st.X < len(data[0])-1 {
				todo = append(todo, step{Y: st.Y, X: st.X + 1, From: FromL, Count: 1, Steps: st.Steps + data[st.Y][st.X+1].Val})
			}
		case FromB:
			// straight
			if st.Count < 3 && st.Y > 0 {
				todo = append(todo, step{Y: st.Y - 1, X: st.X, From: FromB, Count: st.Count + 1, Steps: st.Steps + data[st.Y-1][st.X].Val})
			}
			// left
			if st.X > 0 {
				todo = append(todo, step{Y: st.Y, X: st.X - 1, From: FromR, Count: 1, Steps: st.Steps + data[st.Y][st.X-1].Val})
			}
			// right
			if st.X < len(data[0])-1 {
				todo = append(todo, step{Y: st.Y, X: st.X + 1, From: FromL, Count: 1, Steps: st.Steps + data[st.Y][st.X+1].Val})
			}
		}
	}
	fmt.Println(minEnd)
	writeData(data)
}

func writeData(data [][]cell) {
	out := ""
	for _, row := range data {
		for _, c := range row {
			out += fmt.Sprintf("%5d", c.Steps)
		}
		out += "\n"
	}

	os.WriteFile("./result.txt", []byte(out), 0644)
}
