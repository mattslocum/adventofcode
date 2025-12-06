package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello, playground")
}

var romanNumerals = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

const maxRoman = 3999

func Roman(n int) (string, error) {
	err := validate(n)
	if err != nil {
		return "", err
	}

	out := ""
	for n > 0 {
		digit, unit := parseDigitAndUnit(n)
		out += getSymbol(digit, unit)
		// TODO: Would doing string processing be better than numbers?
		n -= digit * unit
	}
	return out, nil
}

func validate(n int) error {
	if n < 1 {
		return fmt.Errorf("must be a positive number")
	}
	if n > maxRoman {
		return fmt.Errorf("%d is the maximum number allowed", maxRoman)
	}
	return nil
}

func parseDigitAndUnit(n int) (int, int) {
	str := fmt.Sprintf("%d", n)
	digit := int(str[0] - '0')
	tensUnit := generateNumberWithNZeros(len(str) - 1)
	return digit, tensUnit
}

func generateNumberWithNZeros(n int) int {
	return int(math.Pow(float64(10), float64(n)))
}

func getSymbol(digit int, unit int) string {
	one := romanNumerals[unit]
	five := romanNumerals[unit*5]
	ten := romanNumerals[unit*10]
	switch {
	case digit == 1:
		return one
	case digit == 2:
		return one + one
	case digit == 3:
		return one + one + one
	case digit == 4:
		return one + five
	case digit == 5:
		return five
	case digit == 6:
		return five + one
	case digit == 7:
		return five + one + one
	case digit == 8:
		return five + one + one + one
	case digit == 9:
		return one + ten
	}
	return ""
}
