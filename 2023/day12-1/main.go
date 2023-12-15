package main

import (
	"fmt"
	"strings"

	"github.com/merijnf/advent-of-code-go/utilities"
)

func main() {
	lines := utilities.LoadInputStringLines(2023, 12)
	result := 0
	for _, line := range lines {
		splitPattern := strings.Split(line, " ")
		pattern := splitPattern[0]
		groups := utilities.ConvertStrArrToIntArr(strings.Split(splitPattern[1], ","))
		result += calculatePossibilities(pattern, groups)
	}
	fmt.Println(result)
}

func calculatePossibilities(pattern string, groups []int) int {
	if len(groups) == 0 {
		if !strings.Contains(pattern, "#") {
			return 1
		}
		return 0
	}

	if len(pattern) == 0 {
		return 0
	}

	character := pattern[0]
	group := groups[0]
	result := 0

	if character == '.' || character == '?' {
		result += calculatePossibilities(pattern[1:], groups)
	}

	if character == '#' || character == '?' {
		if len(pattern) >= group {
			groupString := pattern[:group]
			groupString = strings.Replace(groupString, "?", "#", -1)
			if !strings.Contains(groupString, ".") {
				if len(pattern) == group {
					if len(groups) == 1 {
						result += 1
					}
				} else {
					characterAfterGroup := pattern[group]
					if characterAfterGroup == '.' || characterAfterGroup == '?' {
						result += calculatePossibilities(pattern[group+1:], groups[1:])
					}
				}
			}
		}
	}

	return result
}
