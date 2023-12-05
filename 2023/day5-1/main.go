package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"

	"github.com/merijnf/advent-of-code-go/utilities"
)

type mapRange struct {
	destRangeStart   int
	sourceRangeStart int
	rangeLength      int
}

func main() {
	lines := utilities.LoadInputStringLines(2023, 5)
	seeds := parseSeeds(lines)
	maps := parseMaps(lines)

	smallest := math.MaxInt

	for _, s := range seeds {
		value := s
		for _, m := range maps {
			for _, r := range m {
				if value >= r.sourceRangeStart && value < r.sourceRangeStart+r.rangeLength {
					value = value - r.sourceRangeStart + r.destRangeStart
					break
				}
			}
		}
		if value < smallest {
			smallest = value
		}
	}
	fmt.Println("final value: " + strconv.Itoa(smallest))
}

func parseSeeds(lines []string) []int {
	//first line is seeds
	seedsLine := lines[0]
	//strip label "seeds:"
	return utilities.ConvertStrArrToIntArr(strings.Fields(seedsLine[6:]))
}

func parseMaps(lines []string) [][]mapRange {
	maps := [][]mapRange{{}}
	mapIndex := 0

	for _, line := range lines[2:] {
		if line == "" {
			mapIndex = mapIndex + 1
			maps = append(maps, []mapRange{})
		} else if unicode.IsLetter(rune(line[0])) {
			// do nothing
		} else {
			// range line
			numbers := utilities.ConvertStrArrToIntArr(strings.Fields(line))
			maps[mapIndex] = append(maps[mapIndex], mapRange{destRangeStart: numbers[0], sourceRangeStart: numbers[1], rangeLength: numbers[2]})
		}
	}
	return maps
}
