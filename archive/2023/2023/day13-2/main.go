package main

import (
	"fmt"
	"slices"

	"github.com/merijnf/advent-of-code-go/utilities"
)

type grid struct {
	verticalLines   []string
	horizontalLines []string
}

func main() {
	lines := utilities.LoadInputStringLines(2023, 13)

	grids := parseGrids(lines)
	result := 0
	for _, grid := range grids {
		result = result + (100 * findReflectionIndex(grid.horizontalLines))
		result = result + findReflectionIndex(grid.verticalLines)
	}

	fmt.Println(result)
}

func parseGrids(lines []string) []grid {
	grids := []grid{}
	currentGrid := grid{horizontalLines: make([]string, 0), verticalLines: make([]string, 0)}
	currentGridIndex := 0
	for _, line := range lines {
		if line == "" {
			grids = append(grids, currentGrid)
			currentGrid = grid{horizontalLines: make([]string, 0), verticalLines: make([]string, 0)}
			currentGridIndex = 0
		} else {
			currentGrid.horizontalLines = append(currentGrid.horizontalLines, line)

			initVerticalLines := false
			if len(currentGrid.verticalLines) == 0 {
				initVerticalLines = true
			}

			for i, character := range line {
				if initVerticalLines {
					currentGrid.verticalLines = append(currentGrid.verticalLines, "")
				}
				currentGrid.verticalLines[i] = currentGrid.verticalLines[i] + string(character)
			}
			currentGridIndex++
		}
	}
	grids = append(grids, currentGrid)
	return grids
}

func findReflectionIndex(lines []string) int {
	for i := range lines {
		linesA := []string{}
		linesB := []string{}
		if i > 0 && i <= len(lines)/2 {
			linesA = lines[:i]
			linesB = slices.Clone(lines[i : i+i])
		}
		if i > len(lines)/2 {
			linesA = lines[i-(len(lines)-i) : i]
			linesB = slices.Clone(lines[i:])
		}

		slices.Reverse(linesB)

		if len(linesA) > 0 && len(linesB) > 0 {

			if differentCharacters(concatStrings(linesA), concatStrings(linesB)) == 1 {
				return i
			}
		}
	}
	return 0
}

func concatStrings(lines []string) string {
	value := ""
	for _, line := range lines {
		value = value + line
	}
	return value
}

func differentCharacters(a string, b string) int {
	count := 0
	for i, char := range a {
		if char != rune(b[i]) {
			count++
		}
	}
	return count
}
