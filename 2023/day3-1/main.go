package main

import (
	"fmt"
	"os"
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
	dat, err := os.ReadFile("../../input/2023/day3/input.txt")
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
		currentPartNumber := ""
		startOfPartNum := -1
		endOfPartNum := -1
		for xIndex := range line {
			token := grid[yIndex][xIndex]
			if token >= '0' && token <= '9' {
				currentPartNumber = currentPartNumber + string(token)
				if startOfPartNum == -1 {
					//new number
					startOfPartNum = xIndex
				}
				if xIndex == len(line)-1 {
					if currentPartNumber == "404" {
						fmt.Println("Debug")
					}
					if startOfPartNum > -1 {
						endOfPartNum = xIndex - 1
						valid := checkSymbol(grid, yIndex, startOfPartNum, endOfPartNum)
						if valid {
							num, err := strconv.Atoi(currentPartNumber)
							check(err)
							sum = sum + num

						}

						fmt.Println(currentPartNumber + " is valid " + strconv.FormatBool(valid))
						currentPartNumber = ""
						startOfPartNum = -1
						endOfPartNum = -1
					}
				}
			} else {
				if startOfPartNum > -1 {
					if currentPartNumber == "404" {
						fmt.Println("Debug")
					}
					endOfPartNum = xIndex - 1
					valid := checkSymbol(grid, yIndex, startOfPartNum, endOfPartNum)
					if valid {
						num, err := strconv.Atoi(currentPartNumber)
						check(err)
						sum = sum + num

					}

					fmt.Println(currentPartNumber + " is valid " + strconv.FormatBool(valid))
					currentPartNumber = ""
					startOfPartNum = -1
					endOfPartNum = -1
				}
			}
		}
	}
	fmt.Println("final value: " + strconv.Itoa(sum))
}

func checkSymbol(grid [][]rune, y int, xStart int, xEnd int) bool {
	yFrom := y - 1
	if yFrom < 0 {
		yFrom = 0
	}
	yTo := y + 1
	if yTo >= len(grid) {
		yTo = len(grid) - 1
	}
	xFrom := xStart - 1
	if xFrom < 0 {
		xFrom = 0
	}
	xTo := xEnd + 1
	if xTo >= len(grid[0]) {
		xTo = len(grid[0]) - 1
	}
	//fmt.Println("checking x from " + strconv.Itoa(xFrom) + " to " + strconv.Itoa(xTo))
	//fmt.Println("checking y from " + strconv.Itoa(yFrom) + " to " + strconv.Itoa(yTo))
	// checked := make([]rune, 0)
	// checkedSting := ""
	for i := yFrom; i <= yTo; i++ {
		for j := xFrom; j <= xTo; j++ {
			value := grid[i][j]
			// checked = append(checked, value)
			// checkedSting = string(checked)
			if value != '.' && value != '\n' {
				if value < '0' || value > '9' {
					// fmt.Println("checked " + checkedSting)
					return true
				}
			}
		}
	}
	// fmt.Println("checked " + checkedSting)
	return false
}

func printGrid(grid [][]rune) {
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
