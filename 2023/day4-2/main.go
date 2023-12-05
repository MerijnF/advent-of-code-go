package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/merijnf/advent-of-code-go/utility"
)

type card struct {
	id                 int
	numbers            []int
	winningNumbers     []int
	winningNumberCount int
	count              int
}

func main() {
	inputLines := utility.LoadInputStringLines(2023, 4)
	cards := make([]card, len(inputLines))
	for i, line := range inputLines {
		cards[i] = parseCard(line)
	}

	for i, card := range cards {
		for j := 1; j <= card.winningNumberCount; j++ {
			index := i + j
			if index > len(cards)-1 {
				break
			}
			cards[i+j].count = cards[i+j].count + 1*card.count
		}
	}
	sum := 0
	for _, card := range cards {
		sum = sum + card.count
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
	utility.CheckError(error)
	winningNumbersStrArr := strings.Split(winningNumbersStr, " ")
	numberStrArr := strings.Split(numbersStr, " ")
	winningNumbersStrArr = slices.DeleteFunc(winningNumbersStrArr, isEmpty)
	numberStrArr = slices.DeleteFunc(numberStrArr, isEmpty)
	winningNumbersArr := convertAllToInt(winningNumbersStrArr)
	numbersArr := convertAllToInt(numberStrArr)
	card := card{id: id, winningNumbers: winningNumbersArr, numbers: numbersArr, count: 1}
	card.winningNumberCount = card.CalculateWinningNumberCount()
	return card
}

func (card card) CalculateWinningNumberCount() int {
	value := 0

	for _, winningNumber := range card.winningNumbers {
		for _, number := range card.numbers {
			if number == winningNumber {
				value = value + 1
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
		utility.CheckError(error)
		numbers[i] = number
	}
	return numbers
}
