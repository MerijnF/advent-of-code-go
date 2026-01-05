package d2

import (
	"strconv"
	"strings"

	"github.com/merijnf/advent-of-code-go/pkg/solver"
)

func Part1(ctx solver.PuzzleContext) (string, error) {
	productRanges := strings.Split(ctx.Input, ",")

	idSum := 0

	for _, productRange := range productRanges {
		ids := strings.Split(productRange, "-")
		firstId := strings.TrimSpace(ids[0])
		lastId := strings.TrimSpace(ids[1])
		firstIdNum, err := strconv.Atoi(firstId)
		if err != nil {
			panic(err)
		}
		lastIdNum, err := strconv.Atoi(lastId)
		if err != nil {
			panic(err)
		}

		for id := firstIdNum; id <= lastIdNum; id++ {
			idStr := strconv.Itoa(id)
			length := len(idStr)
			if length%2 != 0 {
				continue // must be even length
			}
			firstHalf := idStr[:length/2]
			computed := strings.Repeat(firstHalf, 2)
			if computed == idStr {
				idSum += id
			}
		}
	}
	return strconv.Itoa(idSum), nil

}

func Part2(ctx solver.PuzzleContext) (string, error) {
	productRanges := strings.Split(ctx.Input, ",")

	idSum := 0

	for _, productRange := range productRanges {
		ids := strings.Split(productRange, "-")
		firstId := strings.TrimSpace(ids[0])
		lastId := strings.TrimSpace(ids[1])
		firstIdNum, err := strconv.Atoi(firstId)
		if err != nil {
			panic(err)
		}
		lastIdNum, err := strconv.Atoi(lastId)
		if err != nil {
			panic(err)
		}

		for id := firstIdNum; id <= lastIdNum; id++ {
			idStr := strconv.Itoa(id)
			length := len(idStr)
			for divisor := 2; divisor <= length; divisor += 1 {
				if length%divisor != 0 {
					continue // must be evenly divisible
				}
				firstPart := idStr[:length/divisor]
				computed := strings.Repeat(firstPart, divisor)
				if computed == idStr {
					idSum += id
					break
				}
			}
		}
	}
	return strconv.Itoa(idSum), nil
}
