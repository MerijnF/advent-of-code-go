package d3

import (
	"strconv"
	"strings"

	"github.com/merijnf/advent-of-code-go/pkg/solver"
)

func Part1(ctx solver.PuzzleContext) (string, error) {
	banks := strings.Split(ctx.Input, "\r\n")

	solution := 0
	for _, bank := range banks {
		highestJolt := 0
		highestAfterJolt := 0
		batteryCount := len(bank)
		for i, battery := range bank {
			jolt, err := strconv.Atoi(string(battery))
			if err != nil {
				return "", err
			}
			if jolt > highestJolt && i < batteryCount-1 {
				highestJolt = jolt
				highestAfterJolt = 0
			} else {
				if jolt > highestAfterJolt {
					highestAfterJolt = jolt
				}
			}
		}
		highestBattery := strconv.Itoa(highestJolt)
		afterHighestBattery := strconv.Itoa(highestAfterJolt)
		joltValue, err := strconv.Atoi(highestBattery + afterHighestBattery)
		if err != nil {
			return "", err
		}
		solution += joltValue
	}
	return strconv.Itoa(solution), nil
}

func Part2(ctx solver.PuzzleContext) (string, error) {
	banks := strings.Split(ctx.Input, "\r\n")

	solution := 0
	for _, bank := range banks {
		joltCount := 12
		joltValues := make([]int, joltCount)
		batteryCount := len(bank)
		for i, battery := range bank {
			jolt, err := strconv.Atoi(string(battery))
			if err != nil {
				return "", err
			}
			used := false
			for j := 0; j < joltCount; j++ {
				if !used && jolt > joltValues[j] && i < batteryCount-(joltCount-j-1) {
					joltValues[j] = jolt
					used = true
				} else {
					if used {
						// reset all after
						joltValues[j] = 0
					}
				}
			}
		}
		var bankValue strings.Builder
		for _, joltValue := range joltValues {
			bankValue.WriteString(strconv.Itoa(joltValue))
		}
		joltValue, err := strconv.Atoi(bankValue.String())
		if err != nil {
			return "", err
		}
		solution += joltValue
	}
	return strconv.Itoa(solution), nil
}
