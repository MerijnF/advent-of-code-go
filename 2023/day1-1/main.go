package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	dat, err := os.ReadFile("../../input/2023/day1/input.txt")
	check(err)
	datString := string(dat)
	lines := strings.Split(datString, "\n")
	sum := 0
	for _, line := range lines {
		value, err := parseLine((line))
		check(err)
		valueInt, err := strconv.Atoi(value)
		check(err)
		sum = sum + valueInt
	}
	fmt.Println("final value: " + strconv.Itoa(sum))
}

func parseLine(line string) (value string, err error) {
	matchNumbers := regexp.MustCompile(`[0-9]`)
	matches := matchNumbers.FindAllString(line, -1)
	if matches != nil {
		value = matches[0] + matches[len(matches)-1]
		fmt.Println(value)
		return value, nil
	} else {
		return "", errors.New("no value found")
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
