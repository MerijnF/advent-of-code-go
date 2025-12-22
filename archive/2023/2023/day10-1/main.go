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

//const ground = '.'

func main() {
	lines := utilities.LoadInputStringLines(2023, 10)
	pipeMaze := make([][]rune, len(lines))
	for i, line := range lines {
		pipeMaze[i] = []rune(line)
	}
	//find start
	x, y := findStart(pipeMaze)
	prevX, prevY := x, y
	steps := 1
	x, y = step(pipeMaze, x, y, prevX, prevY)
	position := pipeMaze[y][x]
	for position != start {
		nextX, nextY := step(pipeMaze, x, y, prevX, prevY)
		prevX, prevY = x, y
		x, y = nextX, nextY
		position = pipeMaze[y][x]
		steps++
	}
	fmt.Println(steps / 2)
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
