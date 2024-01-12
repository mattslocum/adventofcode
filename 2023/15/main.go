package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type lens struct {
	Str  string
	Hash int
	Opp  string
	Val  int
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lenses := parseInput(string(data))
	boxes := fillBoxes(lenses)
	printLensPower(boxes)
}

func parseInput(data string) []lens {
	strList := strings.Split(string(data), ",")
	lenses := make([]lens, len(strList))
	for i, str := range strList {
		l := lens{}
		if strings.Contains(str, "-") {
			l.Str = str[:len(str)-1]
			l.Hash = hash(l.Str)
			l.Opp = "-"
		} else {
			l.Str = str[:len(str)-2]
			l.Hash = hash(l.Str)
			l.Opp = "="
			n, _ := strconv.Atoi(str[len(str)-1:])
			l.Val = n
		}
		lenses[i] = l
	}

	return lenses
}

func hash(str string) int {
	h := 0
	for _, c := range str {
		h += int(c)
		h *= 17
		h %= 256
	}
	return h
}

func fillBoxes(lenses []lens) [][]lens {
	// TODO: map would be faster than array, but then you lose the order so going with less efficient for now
	boxes := make([][]lens, 256)

	for _, l := range lenses {
		if l.Opp == "=" {
			isFound := false
			for i, box := range boxes[l.Hash] {
				if box.Str == l.Str {
					boxes[l.Hash][i] = l
					isFound = true
					break
				}
			}
			if !isFound {
				boxes[l.Hash] = append(boxes[l.Hash], l)
			}
		} else if l.Opp == "-" {
			newBox := []lens{}
			for _, box := range boxes[l.Hash] {
				if box.Str != l.Str {
					newBox = append(newBox, box)
				}
			}
			boxes[l.Hash] = newBox
		}
	}

	return boxes
}

func printLensPower(boxes [][]lens) {
	total := 0
	for i, box := range boxes {
		for x, l := range box {
			total += (i + 1) * (x + 1) * l.Val
		}
	}
	fmt.Println(total)
}
