package utilities

import (
	"strconv"
)

func ConvertStrArrToIntArr(textArr []string) []int {
	numbers := make([]int, len(textArr))
	for i, text := range textArr {
		number, error := strconv.Atoi(text)
		CheckError(error)
		numbers[i] = number
	}
	return numbers
}
