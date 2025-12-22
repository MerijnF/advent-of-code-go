package main

import (
	"fmt"

	"github.com/merijnf/advent-of-code-go/internal/y2025d01"
)

type DayPart struct {
	day  int
	part int
}

func main() {
	dayPart, testMode := getDayPart()
	var result string
	switch dayPart {
	case DayPart{1, 1}:
		result = y2025d01.Part1(testMode)
	case DayPart{1, 2}:
		result = y2025d01.Part2(testMode)
	default:
		result = "Not implemented"
	}

	fmt.Println("Solution:", result)
}

func getDayPart() (DayPart, bool) {
	fmt.Println("Enter day number (1-31):")
	var day int
	_, err := fmt.Scan(&day)
	if err != nil {
		panic(err)
	}
	fmt.Println("Enter part number (1-2):")
	var part int
	_, err = fmt.Scan(&part)
	if err != nil {
		panic(err)
	}
	fmt.Println("Solving day", day, "part", part)
	fmt.Println("Test mode? (y/n):")
	var testMode bool
	var input string
	_, err = fmt.Scan(&input)
	if err != nil {
		panic(err)
	}
	if input == "y" {
		testMode = true
	}
	return DayPart{day, part}, testMode
}
