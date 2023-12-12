package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("loading input data")
	//dat, err := os.ReadFile("input-test.txt")
	dat, err := os.ReadFile("../../input/2023/day03/input.txt")
	check(err)
	dataString := string(dat)
	lines := strings.Split(strings.ReplaceAll(dataString, "\r\n", "\n"), "\n")
	fmt.Println("creating grid")
	yLen := len(lines)
	grid := make([][]rune, yLen)
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	sum := 0
	printGrid(grid)
	for yIndex, line := range grid {
		for xIndex := range line {
			token := grid[yIndex][xIndex]
			if token == '*' {
				parts := findParts(grid, xIndex, yIndex)
				if len(parts) == 2 {
					value1, err := strconv.Atoi(parts[0])
					check(err)
					value2, err := strconv.Atoi(parts[1])
					check(err)
					value := value1 * value2
					sum = sum + value
				}
			}
		}
	}

	fmt.Println("final value: " + strconv.Itoa(sum))
}

func findParts(grid [][]rune, x int, y int) []string {
	partsFound := make([]string, 0)
	yFrom := y - 1
	if yFrom < 0 {
		yFrom = 0
	}
	yTo := y + 1
	if yTo >= len(grid) {
		yTo = len(grid) - 1
	}
	xFrom := x - 1
	if xFrom < 0 {
		xFrom = 0
	}
	xTo := x + 1
	if xTo >= len(grid[0]) {
		xTo = len(grid[0]) - 1
	}
	for i := yFrom; i <= yTo; i++ {
		for j := xFrom; j <= xTo; j++ {
			value := grid[i][j]
			if value >= '0' && value <= '9' {
				part := ""
				search := value
				left := 0
				right := 0
				for search >= '0' && search <= '9' {
					left = left + 1
					xCheck := j - left
					if xCheck < 0 {
						break
					}
					search = grid[i][xCheck]
				}
				left = left - 1
				search = grid[i][j-left]
				for search >= '0' && search <= '9' {
					part = part + string(search)
					right = right + 1
					xCheck := j - left + right
					if xCheck >= len(grid[0]) {
						break
					}
					search = grid[i][xCheck]
				}
				if !slices.Contains(partsFound, part) {
					partsFound = append(partsFound, part)
				}
			}
		}
	}

	return partsFound
}

func printGrid(grid [][]rune) {
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
