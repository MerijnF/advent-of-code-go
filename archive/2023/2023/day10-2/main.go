package main

import (
	"fmt"

	"github.com/merijnf/advent-of-code-go/utilities"
)

const start = 'S'
const vertical = '|'
const horizontal = '-'
const northEast = 'L'
const northWest = 'J'
const southWest = '7'
const southEast = 'F'

const ground = '.'

func main() {
	lines := utilities.LoadInputStringLines(2023, 10)
	pipeMaze := make([][]rune, len(lines))
	maskedMaze := make([][]rune, len(lines))
	for i, line := range lines {
		pipeMaze[i] = []rune(line)
		emptyRow := make([]rune, len(line))
		for i := range emptyRow {
			emptyRow[i] = ground
		}
		maskedMaze[i] = emptyRow
	}

	//find start
	x, y := findStart(pipeMaze)
	prevX, prevY := x, y
	x, y = step(pipeMaze, x, y, prevX, prevY)
	position := pipeMaze[y][x]
	maskedMaze[y][x] = position
	for position != start {
		nextX, nextY := step(pipeMaze, x, y, prevX, prevY)
		prevX, prevY = x, y
		x, y = nextX, nextY
		position = pipeMaze[y][x]
		maskedMaze[y][x] = position
	}

	maskedMaze[y][x] = starTile(maskedMaze, x, y)

	result := 0
	for y, row := range maskedMaze {
		for x := range row {
			if inPipe(maskedMaze, x, y) {
				result++
				maskedMaze[y][x] = '#'
			}
		}
	}
	for _, row := range maskedMaze {
		fmt.Println(string(row))
	}
	fmt.Println(result)
}

func inPipe(maze [][]rune, x, y int) bool {
	position := maze[y][x]
	if position != ground {
		return false
	}
	crossings := 0
	lastCrossing := ground
	for i := x; i >= 0; i-- {
		check := maze[y][i]
		switch check {
		case vertical:
			crossings++
			lastCrossing = check
		case northWest:
			crossings++
			lastCrossing = check
		case southWest:
			crossings++
			lastCrossing = check
		case northEast:
			if lastCrossing == northWest {
				crossings--
			}
		case southEast:
			if lastCrossing == southWest {
				crossings--
			}
		}
	}
	return crossings%2 != 0
}

func starTile(maze [][]rune, x, y int) rune {
	north := false
	east := false
	south := false
	west := false
	if y > 0 {
		checking := maze[y-1][x]
		north = checking == vertical || checking == southWest || checking == southEast
	}
	if x < len(maze[y]) {
		checking := maze[y][x+1]
		east = checking == horizontal || checking == northWest || checking == southWest
	}
	if y < len(maze) {
		checking := maze[y+1][x]
		south = checking == vertical || checking == northEast || checking == northWest
	}
	if x > 0 {
		checking := maze[y][x-1]
		west = checking == horizontal || checking == northEast || checking == southEast
	}

	if north && south {
		return vertical
	}
	if west && east {
		return horizontal
	}
	if south && west {
		return southWest
	}
	if south && east {
		return southEast
	}
	if north && east {
		return northEast
	}
	if north && west {
		return northWest
	}
	return start
}

func findStart(maze [][]rune) (x, y int) {
	for y, row := range maze {
		for x, position := range row {
			if position == start {
				return x, y
			}
		}
	}
	return -1, -1
}
func step(maze [][]rune, x, y, prevX, prevY int) (nextX, nextY int) {
	switch maze[y][x] {
	case start:
		if x != 0 {
			checking := maze[y][x-1]
			if checking == horizontal || checking == northEast || checking == southEast {
				return x - 1, y
			}
		}
		if x != len(maze[0]) {
			checking := maze[y][x+1]
			if checking == horizontal || checking == southWest || checking == northWest {
				return x + 1, y
			}
		}
		if y != 0 {
			checking := maze[y-1][x]
			if checking == vertical || checking == southWest || checking == southEast {
				return x, y - 1
			}
		}
		if y != len(maze) {
			checking := maze[y+1][x]
			if checking == vertical || checking == northEast || checking == northWest {
				return x, y + 1
			}
		}
	case vertical:
		nextX, nextY = x, y+1
		if nextX == prevX && nextY == prevY {
			nextX, nextY = x, y-1
		}
		return nextX, nextY
	case horizontal:
		nextX, nextY = x+1, y
		if nextX == prevX && nextY == prevY {
			nextX, nextY = x-1, y
		}
		return nextX, nextY
	case northEast:
		nextX, nextY = x, y-1
		if nextX == prevX && nextY == prevY {
			nextX, nextY = x+1, y
		}
		return nextX, nextY
	case northWest:
		nextX, nextY = x, y-1
		if nextX == prevX && nextY == prevY {
			nextX, nextY = x-1, y
		}
		return nextX, nextY
	case southWest:
		nextX, nextY = x, y+1
		if nextX == prevX && nextY == prevY {
			nextX, nextY = x-1, y
		}
		return nextX, nextY
	case southEast:
		nextX, nextY = x, y+1
		if nextX == prevX && nextY == prevY {
			nextX, nextY = x+1, y
		}
		return nextX, nextY
	}
	return -1, -1
}
