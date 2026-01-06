package d3

import (
	"strconv"
	"strings"

	"github.com/merijnf/advent-of-code-go/pkg/solver"
	"github.com/merijnf/advent-of-code-go/pkg/utils/strutil"
)

func Part1(ctx solver.PuzzleContext) (string, error) {
	return solve(ctx.Input, 2)
}

func Part2(ctx solver.PuzzleContext) (string, error) {
	return solve(ctx.Input, 12)
}

func solve(input string, joltCount int) (string, error) {
	banks := strutil.SplitLinesAndNormalizeSeq(input, false)

	solution := 0
	for bank := range banks {
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
