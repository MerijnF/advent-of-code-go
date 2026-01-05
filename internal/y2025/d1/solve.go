package d1

import (
	"strconv"
	"strings"

	"github.com/merijnf/advent-of-code-go/pkg/solver"
)

const left = "L"
const right = "R"

func Part1(ctx solver.PuzzleContext) (string, error) {
	lines := strings.Split(ctx.Input, "\r\n")
	position := 50
	atZero := 0
	for _, line := range lines {
		if strings.HasPrefix(line, left) {
			move, err := strconv.Atoi(strings.TrimPrefix(line, left))
			if err != nil {
				panic(err)
			}
			position -= move
		} else if strings.HasPrefix(line, right) {
			move, err := strconv.Atoi(strings.TrimPrefix(line, right))
			if err != nil {
				panic(err)
			}
			position += move
		} else {
			panic("Unknown direction")
		}
		for position < 0 {
			position += 100
		}
		for position >= 100 {
			position -= 100
		}
		if position == 0 {
			atZero++
		}
	}

	return strconv.Itoa(atZero), nil
}

func Part2(ctx solver.PuzzleContext) (string, error) {
	lines := strings.Split(ctx.Input, "\r\n")
	position := 50
	zeroCount := 0
	atZero := false
	for _, line := range lines {
		if strings.HasPrefix(line, left) {
			move, err := strconv.Atoi(strings.TrimPrefix(line, left))
			if err != nil {
				panic(err)
			}
			position -= move
		} else if strings.HasPrefix(line, right) {
			move, err := strconv.Atoi(strings.TrimPrefix(line, right))
			if err != nil {
				panic(err)
			}
			position += move
		} else {
			panic("Unknown direction")
		}
		for position < 0 {
			if !atZero {
				zeroCount++
			}
			if atZero && position < -100 {
				// if we start at zero and move more than 100 left, count the zero crossing
				// if we are less than 100 away from zero we just moved left and did not pass zero
				zeroCount++
			}
			position += 100
		}
		for position >= 100 {
			if position != 100 {
				zeroCount++
			}
			position -= 100
		}
		if position == 0 {
			zeroCount++
			atZero = true
		} else {
			atZero = false
		}
	}

	return strconv.Itoa(zeroCount), nil
}
