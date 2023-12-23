package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pos struct {
	X     int
	Y     int
	PrevX int
	PrevY int
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	grid := parseInput(string(data))

	startY, startX := findStart(&grid)

	dist := findDist(&grid, startY, startX)
	fmt.Println(dist / 2)
	renderResult(grid, dist)
}

func parseInput(data string) [][]string {
	lines := strings.Split(string(data), "\n")
	result := make([][]string, len(lines))
	for lineIdx, line := range lines {
		result[lineIdx] = strings.Split(string(line), "")
	}
	return result
}

func findStart(grid *[][]string) (int, int) {
	var (
		startY int
		startX int
	)
	lenGridY := len(*grid)
	lenGridX := len((*grid)[0])
	for y := 0; startY == 0 && y < lenGridY; y++ {
		for x := 0; x < lenGridX; x++ {
			if (*grid)[y][x] == "S" {
				startY = y
				startX = x
				break
			}
		}
	}
	return startY, startX
}

func findDist(grid *[][]string, startY int, startX int) int {
	walk := setupWalker(grid, startY, startX)
	steps := 1
	for ; (*grid)[walk.Y][walk.X][0] != 'S'; steps++ {
		walk = step(grid, walk, steps)
	}

	return steps
}

func setupWalker(grid *[][]string, startY int, startX int) (w pos) {
	w = pos{
		PrevX: startX,
		PrevY: startY,
	}

	up := (*grid)[startY-1][startX]
	if up == "|" || up == "7" || up == "F" {
		w.X = startX
		w.Y = startY - 1
		return
	}

	right := (*grid)[startY][startX+1]
	if right == "-" || right == "7" || right == "J" {
		w.X = startX + 1
		w.Y = startY
		return
	}

	down := (*grid)[startY+1][startX]
	if down == "|" || down == "L" || down == "J" {
		w.X = startX
		w.Y = startY + 1
		return
	}

	// only a problem for test data
	if startX > 0 {
		left := (*grid)[startY][startX-1]
		if left == "-" || left == "L" || left == "F" {
			w.X = startX - 1
			w.Y = startY
			return
		}
	}
	return
}

// not going to worry about out of bounds
func step(grid *[][]string, cur pos, steps int) pos {
	/**
	  | is a vertical pipe connecting north and south.
	  - is a horizontal pipe connecting east and west.
	  L is a 90-degree bend connecting north and east.
	  J is a 90-degree bend connecting north and west.
	  7 is a 90-degree bend connecting south and west.
	  F is a 90-degree bend connecting south and east.
	  . is ground; there is no pipe in this tile.
	  S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
	*/
	next := pos{
		X:     cur.X,
		Y:     cur.Y,
		PrevX: cur.X,
		PrevY: cur.Y,
	}
	switch (*grid)[cur.Y][cur.X] {
	case "|":
		if cur.PrevY < cur.Y {
			next.Y++
		} else {
			next.Y--
		}
		next.X = cur.X
	case "-":
		if cur.PrevX < cur.X {
			next.X++
		} else {
			next.X--
		}
		next.Y = cur.Y
	case "L":
		if cur.PrevY != cur.Y {
			next.Y = cur.Y
			next.X++
		} else {
			next.X = cur.X
			next.Y--
		}
	case "J":
		if cur.PrevY != cur.Y {
			next.Y = cur.Y
			next.X--
		} else {
			next.X = cur.X
			next.Y--
		}
	case "7":
		if cur.PrevY != cur.Y {
			next.Y = cur.Y
			next.X--
		} else {
			next.X = cur.X
			next.Y++
		}
	case "F":
		if cur.PrevY != cur.Y {
			next.Y = cur.Y
			next.X++
		} else {
			next.X = cur.X
			next.Y++
		}
	}
	(*grid)[cur.Y][cur.X] = (*grid)[cur.Y][cur.X] + strconv.Itoa(steps)
	return next
}

func renderResult(grid [][]string, dist int) {
	// TODO: Consider golang templates
	body := ""
	for _, row := range grid {
		line := "<div class='row'>"
		for _, cell := range row {
			if cell[0] == 'S' {
				line += "<span class='s'>" + cell[:1] + "</span>"
			} else if len(cell) > 1 {
				i, _ := strconv.Atoi(cell[1:])
				h := i / (dist / 256)
				line += fmt.Sprintf("<span class='v' style='background: hsl(%d, 100%%, 30%%)'>%s</span>", h, cell[:1])
			} else if cell == "." {
				line += cell
			} else {
				line += " "
			}
		}
		body += line + "</div>"
	}
	html := `<!DOCTYPE html>
<html lang="en">
<style>
.row {
	letter-spacing: 7px;
}
span.s {
	background: #00ff51;
	color: black;
}
span.v {
	background: #9c0000;
	color: white;
}
</style>
<pre>%s</pre>
</html>
	`
	html = fmt.Sprintf(html, body)

	os.WriteFile("./result.html", []byte(html), 0644)
}
