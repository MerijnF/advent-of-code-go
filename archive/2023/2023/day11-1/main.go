package main

import (
	"fmt"

	"github.com/merijnf/advent-of-code-go/utilities"
)

type galaxy struct {
	x int
	y int
}

func main() {
	lines := utilities.LoadInputStringLines(2023, 11)
	galaxies, emptyRows, emptyColumns := parseInput(lines)

	result := 0

	for i, galaxyA := range galaxies {
		for j, galaxyB := range galaxies {
			if j > i {
				distance := abs(galaxyA.x-galaxyB.x) + abs(galaxyA.y-galaxyB.y)
				for i := min(galaxyA.x, galaxyB.x); i < max(galaxyA.x, galaxyB.x); i++ {
					if emptyColumns[i] {
						distance++
					}
				}
				for i := min(galaxyA.y, galaxyB.y); i < max(galaxyA.y, galaxyB.y); i++ {
					if emptyRows[i] {
						distance++
					}
				}
				result = result + distance
			}
		}
	}

	fmt.Println(result)
}

func parseInput(input []string) (galaxies []galaxy, emptyRows []bool, emptyColumns []bool) {
	emptyColumns = make([]bool, len(input[0]))
	for i := range emptyColumns {
		emptyColumns[i] = true
	}
	emptyRows = make([]bool, len(input))
	for i := range emptyRows {
		emptyRows[i] = true
	}
	galaxies = []galaxy{}

	for y, line := range input {
		for x, token := range line {
			if token == '#' {
				galaxies = append(galaxies, galaxy{x: x, y: y})
				emptyColumns[x] = false
				emptyRows[y] = false
			}
		}
	}

	return
}

func abs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}
