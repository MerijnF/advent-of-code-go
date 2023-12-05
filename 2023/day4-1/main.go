package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/merijnf/advent-of-code-go/utilities"
)

type card struct {
	id             int
	numbers        []int
	winningNumbers []int
}

func main() {
	inputLines := utilities.LoadInputStringLines(2023, 4)
	sum := 0
	cards := make([]card, len(inputLines))
	for i, line := range inputLines {
		cards[i] = parseCard(line)
		sum = sum + cards[i].CalculateCardValue()
	}

	fmt.Println("final value: " + strconv.Itoa(sum))
}

func parseCard(line string) card {
	splitIdNumbers := strings.Split(line, ":")
	idStr := splitIdNumbers[0]
	splitNumbers := strings.Split(splitIdNumbers[1], "|")
	winningNumbersStr := splitNumbers[0]
	numbersStr := splitNumbers[1]
	//parse id
	id, error := strconv.Atoi(strings.Trim(idStr[4:], " "))
	utilities.CheckError(error)
	winningNumbersStrArr := strings.Split(winningNumbersStr, " ")
	numberStrArr := strings.Split(numbersStr, " ")
	winningNumbersStrArr = slices.DeleteFunc(winningNumbersStrArr, isEmpty)
	numberStrArr = slices.DeleteFunc(numberStrArr, isEmpty)
	winningNumbersArr := convertAllToInt(winningNumbersStrArr)
	numbersArr := convertAllToInt(numberStrArr)
	return card{id: id, winningNumbers: winningNumbersArr, numbers: numbersArr}
}

func (card card) CalculateCardValue() int {
	value := 0

	for _, winningNumber := range card.winningNumbers {
		for _, number := range card.numbers {
			if number == winningNumber {
				if value == 0 {
					value = 1
				} else {
					value = value * 2
				}
			}
		}
	}

	return value
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
		utilities.CheckError(error)
		numbers[i] = number
	}
	return numbers
}
