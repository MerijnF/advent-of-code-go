package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/merijnf/advent-of-code-go/utilities"
)

func main() {
	lines := utilities.LoadInputStringLines(2023, 9)
	sum := 0
	for _, line := range lines {
		numbers := parseLines(line)
		extrapolated := extrapolateRange(numbers)
		sum = sum + extrapolated
		fmt.Println(numbers)
		fmt.Println(extrapolated)
	}
	fmt.Println("Final value: " + strconv.Itoa(sum))
}

func parseLines(line string) []int {
	fields := strings.Fields(line)
	numbers := make([]int, len(fields))
	for i, field := range fields {
		var error error
		numbers[i], error = strconv.Atoi(field)
		utilities.CheckError(error)
	}
	return numbers
}

func extrapolateRange(numbers []int) int {
	extrapolated := make([][]int, 1)
	extrapolated[0] = numbers
	notAllZero := true
	for i := 0; notAllZero; i++ {
		notAllZero = false
		level := make([]int, len(extrapolated[i])-1)
		for j, number := range extrapolated[i] {
			if j < len(extrapolated[i])-1 {
				level[j] = extrapolated[i][j+1] - number
				if level[j] != 0 {
					notAllZero = true
				}
			}
		}
		extrapolated = append(extrapolated, level)
	}
	slices.Reverse(extrapolated)
	extrapolatedValue := 0
	for _, level := range extrapolated {
		first := level[0]
		extrapolatedValue = first - extrapolatedValue
	}
	return extrapolatedValue
}
