package main

import (
	"fmt"
	"os"
	"strings"
)

type piece int

const (
	Space piece = iota
	Vertical
	Horizontal
	Positive // slope plane positive
	Negative // slope plane negative
)

type tile struct {
	Visited bool
	Type    piece
	UsedA   bool // split used for straight. top used for vertical.
	UsedB   bool // bottom used for vertical
}

type dir int

const (
	L dir = iota
	R
	U
	D
)

type step struct {
	Dir dir // direction moving
	Y   int
	X   int
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	tiles := parseInput(string(data))
	tiles = runReflections(tiles)
	countVisited(tiles)
}

func parseInput(data string) [][]tile {
	lines := strings.Split(string(data), "\n")
	tiles := make([][]tile, len(lines))
	for y, line := range lines {
		tiles[y] = make([]tile, len(line))
		for x, c := range line {
			switch c {
			case '.':
				tiles[y][x] = tile{}
			case '|':
				tiles[y][x] = tile{Type: Vertical}
			case '-':
				tiles[y][x] = tile{Type: Horizontal}
			case '/':
				tiles[y][x] = tile{Type: Positive}
			case '\\':
				tiles[y][x] = tile{Type: Negative}
			}
		}
	}
	return tiles
}

func runReflections(tiles [][]tile) [][]tile {
	todo := []step{
		{Y: 0, X: 0, Dir: R},
	}

	for len(todo) > 0 {
		st := todo[0]
		todo = todo[1:]
		if st.X < 0 || st.X > len(tiles[0])-1 || st.Y < 0 || st.Y > len(tiles)-1 {
			continue
		}
		cur := &tiles[st.Y][st.X]
		cur.Visited = true

		switch cur.Type {
		case Space:
			switch st.Dir {
			case L:
				todo = append(todo, moveLeft(st))
			case R:
				todo = append(todo, moveRight(st))
			case U:
				todo = append(todo, moveUp(st))
			case D:
				todo = append(todo, moveDown(st))
			}

		case Vertical:
			switch st.Dir {
			case L, R:
				if cur.UsedA {
					continue
				}
				cur.UsedA = true
				todo = append(todo, moveUp(st))
				todo = append(todo, moveDown(st))
			case U:
				todo = append(todo, moveUp(st))
			case D:
				todo = append(todo, moveDown(st))
			}

		case Horizontal:
			switch st.Dir {
			case U, D:
				if cur.UsedA {
					continue
				}
				cur.UsedA = true
				todo = append(todo, moveLeft(st))
				todo = append(todo, moveRight(st))
			case L:
				todo = append(todo, moveLeft(st))
			case R:
				todo = append(todo, moveRight(st))
			}

		case Positive:
			switch st.Dir {
			case L:
				if cur.UsedB {
					continue
				}
				cur.UsedB = true
				todo = append(todo, moveDown(st))
			case R:
				if cur.UsedA {
					continue
				}
				cur.UsedA = true
				todo = append(todo, moveUp(st))
			case U:
				if cur.UsedB {
					continue
				}
				cur.UsedB = true
				todo = append(todo, moveRight(st))
				cur.UsedB = true
			case D:
				if cur.UsedA {
					continue
				}
				cur.UsedA = true
				todo = append(todo, moveLeft(st))
			}

		case Negative:
			switch st.Dir {
			case L:
				if cur.UsedA {
					continue
				}
				cur.UsedA = true
				todo = append(todo, moveUp(st))
			case R:
				if cur.UsedB {
					continue
				}
				cur.UsedB = true
				todo = append(todo, moveDown(st))
			case U:
				if cur.UsedB {
					continue
				}
				cur.UsedB = true
				todo = append(todo, moveLeft(st))
			case D:
				if cur.UsedA {
					continue
				}
				cur.UsedA = true
				todo = append(todo, moveRight(st))
			}
		}
	}
	return tiles
}

func moveLeft(st step) step {
	return step{Y: st.Y, X: st.X - 1, Dir: L}
}
func moveRight(st step) step {
	return step{Y: st.Y, X: st.X + 1, Dir: R}
}
func moveUp(st step) step {
	return step{Y: st.Y - 1, X: st.X, Dir: U}
}
func moveDown(st step) step {
	return step{Y: st.Y + 1, X: st.X, Dir: D}
}

func countVisited(tiles [][]tile) {
	count := 0
	for _, row := range tiles {
		for _, t := range row {
			if t.Visited {
				count++
			}
		}
	}
	fmt.Println(count)
}
