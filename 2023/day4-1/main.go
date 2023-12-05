package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/merijnf/advent-of-code-go/utility"
)

type card struct {
	id             int
	numbers        []int
	winningNumbers []int
}

func main() {
	inputLines := utility.LoadInputStringLines(2023, 4)
	for _, line := range inputLines {
		parseCard(line)
	}
}

func parseCard(line string) card {
	splitIdNumbers := strings.Split(line, ":")
	idStr := splitIdNumbers[0]
	splitNumbers := strings.Split(splitIdNumbers[1], "|")
	winningNumbersStr := splitNumbers[0]
	numbersStr := splitIdNumbers[1]
	//parse id
	id, error := strconv.Atoi(strings.Trim(idStr[4:], " "))
	utility.CheckError(error)
	winningNumbersStrArr := strings.Split(winningNumbersStr, " ")
	numberStrArr := strings.Split(numbersStr, " ")
	winningNumbersStrArr = slices.DeleteFunc(winningNumbersStrArr, isEmpty)
	numberStrArr = slices.DeleteFunc(numberStrArr, isEmpty)
	winningNumbersArr := convertAllToInt(winningNumbersStrArr)
	numbersArr := convertAllToInt(numberStrArr)
	return card{id: id, winningNumbers: winningNumbersArr, numbers: numbersArr}
}

func isEmpty(input string) bool {
	if input == "" {
		return true
	}
	return false
}

func convertAllToInt(textArr []string) []int {
	numbers := make([]int, len(textArr))
	for i, text := range textArr {
		number, error := strconv.Atoi(text)
		utility.CheckError(error)
		numbers[i] = number
	}
	return numbers
}
