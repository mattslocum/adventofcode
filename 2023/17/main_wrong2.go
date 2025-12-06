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
	Val   int
	Steps map[int]int
}

type step struct {
	Count int // distance on current direction
	From  dir
	PrevX int // for convenience
	PrevY int // for convenience
	Steps int
	X     int
	Y     int
}

func main() {
	data, err := os.ReadFile("./test2.txt")
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
			cells[y][x] = cell{Val: i, Steps: map[int]int{}}
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
		// if st.Steps > 100 && st.Y < 7 {
		// fmt.Println(st.Y, st.X, st.Steps)
		// }

		// if st.Y == len(data) -1 && st.X == len(data[0]) -1 {

		// }
		if st.Y == len(data)-1 && st.X == len(data[0])-1 {
			if st.Steps < minEnd {
				minEnd = st.Steps
				fmt.Println("end", total, st, cur, minEnd)
			}
			continue
		}

		if st.Steps > minEnd {
			continue
		}

		// process current item
		if cur.Steps[st.Count] == 0 || st.Steps <= cur.Steps[st.Count] {
			cur.Steps[st.Count] = st.Steps
			// } else if st.Count != 2 {
		} else {
			// fmt.Println("skip", st.Y, st.X)
			continue
		}

		// breadth first search
		switch st.From {
		case FromL:
			// straight
			if st.Count < 3 && st.X < len(data[0])-1 {
				todo = append(todo, step{Y: st.Y, X: st.X + 1, From: FromL, Count: st.Count + 1, PrevX: st.X, PrevY: st.Y, Steps: st.Steps + data[st.Y][st.X+1].Val})
			}
			// up
			if st.Y > 0 {
				todo = append(todo, step{Y: st.Y - 1, X: st.X, From: FromB, Count: 1, PrevX: st.X, PrevY: st.Y, Steps: st.Steps + data[st.Y-1][st.X].Val})
			}
			// down
			if st.Y < len(data)-1 {
				todo = append(todo, step{Y: st.Y + 1, X: st.X, From: FromT, Count: 1, PrevX: st.X, PrevY: st.Y, Steps: st.Steps + data[st.Y+1][st.X].Val})
			}
		case FromR:
			// straight
			if st.Count < 3 && st.X > 0 {
				todo = append(todo, step{Y: st.Y, X: st.X - 1, From: FromR, Count: st.Count + 1, PrevX: st.X, PrevY: st.Y, Steps: st.Steps + data[st.Y][st.X-1].Val})
			}
			// up
			if st.Y > 0 {
				todo = append(todo, step{Y: st.Y - 1, X: st.X, From: FromB, Count: 1, PrevX: st.X, PrevY: st.Y, Steps: st.Steps + data[st.Y-1][st.X].Val})
			}
			// down
			if st.Y < len(data)-1 {
				todo = append(todo, step{Y: st.Y + 1, X: st.X, From: FromT, Count: 1, PrevX: st.X, PrevY: st.Y, Steps: st.Steps + data[st.Y+1][st.X].Val})
			}
		case FromT:
			// straight
			if st.Count < 3 && st.Y < len(data)-1 {
				todo = append(todo, step{Y: st.Y + 1, X: st.X, From: FromT, Count: st.Count + 1, PrevX: st.X, PrevY: st.Y, Steps: st.Steps + data[st.Y+1][st.X].Val})
			}
			// left
			if st.X > 0 {
				todo = append(todo, step{Y: st.Y, X: st.X - 1, From: FromR, Count: 1, PrevX: st.X, PrevY: st.Y, Steps: st.Steps + data[st.Y][st.X-1].Val})
			}
			// right
			if st.X < len(data[0])-1 {
				todo = append(todo, step{Y: st.Y, X: st.X + 1, From: FromL, Count: 1, PrevX: st.X, PrevY: st.Y, Steps: st.Steps + data[st.Y][st.X+1].Val})
			}
		case FromB:
			// straight
			if st.Count < 3 && st.Y > 0 {
				todo = append(todo, step{Y: st.Y - 1, X: st.X, From: FromB, Count: st.Count + 1, PrevX: st.X, PrevY: st.Y, Steps: st.Steps + data[st.Y-1][st.X].Val})
			}
			// left
			if st.X > 0 {
				todo = append(todo, step{Y: st.Y, X: st.X - 1, From: FromR, Count: 1, PrevX: st.X, PrevY: st.Y, Steps: st.Steps + data[st.Y][st.X-1].Val})
			}
			// right
			if st.X < len(data[0])-1 {
				todo = append(todo, step{Y: st.Y, X: st.X + 1, From: FromL, Count: 1, PrevX: st.X, PrevY: st.Y, Steps: st.Steps + data[st.Y][st.X+1].Val})
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
			out += fmt.Sprintf("%5d", minMapVal(c.Steps))
		}
		out += "\n"
	}

	os.WriteFile("./result.txt", []byte(out), 0644)
}

func minMapVal(data map[int]int) int {
	min := 0
	setOnce := false
	for _, val := range data {
		if !setOnce || val < min {
			setOnce = true
			min = val
		}
	}
	return min
}
