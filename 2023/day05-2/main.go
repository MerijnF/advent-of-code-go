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

type seedRange struct {
	rangeStart  int
	rangeLength int
}

func main() {
	fmt.Println(math.MaxInt)
	lines := utilities.LoadInputStringLines(2023, 5)
	seedRanges := parseSeedRanges(lines)
	maps := parseMaps(lines)

	results := make(chan int, len(seedRanges))

	for i, sr := range seedRanges {
		go func(sr seedRange, index int) {
			fmt.Println("starting range " + strconv.Itoa(index))
			rangeEnd := sr.rangeStart + sr.rangeLength
			fmt.Println("range " + strconv.Itoa(index) + " from " + strconv.Itoa(sr.rangeStart) + " to " + strconv.Itoa(rangeEnd) + " length " + strconv.Itoa(sr.rangeLength))
			smallest := math.MaxInt
			for seed := sr.rangeStart; seed < rangeEnd; seed++ {
				value := seed
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
			fmt.Println("ending range " + strconv.Itoa(index))
			fmt.Println("smallest in range " + strconv.Itoa(index) + " is " + strconv.Itoa(smallest))
			results <- smallest
		}(sr, i)
	}
	verySmallest := math.MaxInt
	for i := 1; i <= len(seedRanges); i++ {
		test := <-results
		if test < verySmallest {
			verySmallest = test
		}
	}

	fmt.Println("final value: " + strconv.Itoa(verySmallest))
}

func parseSeedRanges(lines []string) []seedRange {
	//first line is seeds
	seedsLine := lines[0]
	//strip label "seeds:"
	seedNumbers := utilities.ConvertStrArrToIntArr(strings.Fields(seedsLine[6:]))
	seedRanges := make([]seedRange, len(seedNumbers)/2)
	for i := 0; i < len(seedNumbers); i = i + 2 {
		seedRanges[i/2] = seedRange{rangeStart: seedNumbers[i], rangeLength: seedNumbers[i+1]}
	}
	return seedRanges
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
