package main

import (
	"fmt"

	"github.com/merijnf/advent-of-code-go/internal/y2025"
	"github.com/merijnf/advent-of-code-go/pkg/solver"
)

type DayPart struct {
	day  int
	part int
}

func main() {
	// register solvers
	y2025.RegisterSolvers()
	// solve
	dayPart := getDayPart()
	result, err := solver.Solve(2025, dayPart.day, dayPart.part)
	if err != nil {
		panic(err)
	}
	fmt.Println("Solution:", result)
}

func getDayPart() DayPart {
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

	return DayPart{day, part}
}
