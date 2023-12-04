package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	dat, err := os.ReadFile("input.txt")
	check(err)
	datString := string(dat)
	lines := strings.Split(datString, "\n")
	sum := 0
	for _, line := range lines {
		value := parseLine((line))
		valueInt, err := strconv.Atoi(value)
		check(err)
		sum = sum + valueInt
		//fmt.Scanln()
	}
	fmt.Println("final value: " + strconv.Itoa(sum))
}

func parseLine(line string) string {
	fmt.Println(line)
	first := findFirst(line)
	last := findLast(line)
	fmt.Println("first: " + first + " last: " + last)
	firstParsed := parseValue(first)
	lastParsed := parseValue(last)
	fmt.Println("first: " + firstParsed + " last: " + lastParsed)
	return firstParsed + lastParsed
}

func findFirst(line string) string {
	matchNumbers := regexp.MustCompile(`[1-9]|one|two|three|four|five|six|seven|eight|nine`)
	for i := 1; i < len(line); i++ {
		slice := line[:i]
		found := matchNumbers.FindString(slice)
		if found != "" {
			return found
		}
	}
	return ""
}

func findLast(line string) string {
	matchNumbers := regexp.MustCompile(`[1-9]|one|two|three|four|five|six|seven|eight|nine`)
	for i := len(line) - 1; i >= 0; i-- {
		slice := line[i:]
		found := matchNumbers.FindString(slice)
		if found != "" {
			return found
		}
	}
	return ""
}

func parseValue(input string) string {
	matchText := regexp.MustCompile(`one|two|three|four|five|six|seven|eight|nine`)
	if matchText.MatchString(input) {
		switch input {
		case "one":
			return "1"
		case "two":
			return "2"
		case "three":
			return "3"
		case "four":
			return "4"
		case "five":
			return "5"
		case "six":
			return "6"
		case "seven":
			return "7"
		case "eight":
			return "8"
		case "nine":
			return "9"
		}
		return input
	} else {
		return input
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
