package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	Dir   string
	Steps int
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	insts := parseInput(string(data))
	walk(insts)
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

func walk(insts []instruction) {
	var (
		curY, curX, maxY, maxX, minY, minX int
	)
	field := map[int]map[int]string{}
	field[0] = map[int]string{0: "S"}
	for _, inst := range insts {
		switch inst.Dir {
		case "R":
			for i := 1; i <= inst.Steps; i++ {
				field[curY][curX+i] = inst.Dir
			}
			curX += inst.Steps
			if curX > maxX {
				maxX = curX
			}
		case "L":
			for i := 1; i <= inst.Steps; i++ {
				field[curY][curX-i] = inst.Dir
			}
			curX -= inst.Steps
			if curX < minX {
				minX = curX
			}
		case "U":
			field[curY][curX] = inst.Dir // to help the count logic later
			for i := 1; i <= inst.Steps; i++ {
				if _, ok := field[curY-i]; !ok {
					field[curY-i] = map[int]string{}
				}
				field[curY-i][curX] = inst.Dir
			}
			curY -= inst.Steps
			if curY < minY {
				minY = curY
			}
		case "D":
			field[curY][curX] = inst.Dir // to help the count logic later
			for i := 1; i <= inst.Steps; i++ {
				if _, ok := field[curY+i]; !ok {
					field[curY+i] = map[int]string{}
				}
				field[curY+i][curX] = inst.Dir
			}
			curY += inst.Steps
			if curY > maxY {
				maxY = curY
			}
		}
	}
	writeData(field, maxY, maxX, minY, minX)

	filled := 0
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if _, has := field[y][x]; has {
				filled++
			} else {
				// check filled left
				countLeft := 0
				prevDir := ""
				for l := x; l >= minX; l-- {
					if val, has := field[y][l]; has {
						if (val == "U" || val == "D") && val != prevDir {
							countLeft++
							prevDir = val
						}
					} else {
						prevDir = ""
					}
				}
				if countLeft%2 == 1 {
					field[y][x] = "X"
					filled++
				}
			}
		}
	}
	writeData(field, maxY, maxX, minY, minX)
	fmt.Println(filled)
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
