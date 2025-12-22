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

func ConvertIntArrToStrArr(textArr []int) []string {
	strings := make([]string, len(textArr))
	for i, text := range textArr {
		number := strconv.Itoa(text)
		strings[i] = number
	}
	return strings
}
