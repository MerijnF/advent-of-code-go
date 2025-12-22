package y2025d02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/merijnf/advent-of-code-go/pkg/input"
)

const left = "L"
const right = "R"

func Part1(testMode bool) string {
	input := input.ReadInput(2025, 2, testMode)
	productRanges := strings.Split(input, ",")

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
	return strconv.Itoa(idSum)

}

func Part2(testMode bool) string {
	input := input.ReadInput(2025, 2, testMode)
	productRanges := strings.Split(input, ",")

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
		fmt.Println(productRange)

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
					fmt.Println(strconv.Itoa(id))
					idSum += id
					break
				}
			}
		}
	}
	return strconv.Itoa(idSum)
}
