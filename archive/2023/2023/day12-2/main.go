package main

import (
	"fmt"
	"strings"

	"github.com/merijnf/advent-of-code-go/utilities"
)

var cache = make(map[string]int)

func main() {
	lines := utilities.LoadInputStringLines(2023, 12)
	result := 0
	for _, line := range lines {
		splitPattern := strings.Split(line, " ")
		pattern := splitPattern[0]
		unfoldedPattern := ""
		for i := 0; i < 5; i++ {
			unfoldedPattern = unfoldedPattern + pattern
			if i < 4 {
				unfoldedPattern = unfoldedPattern + "?"
			}
		}
		groupsString := splitPattern[1]
		unfoldedGroupsString := ""
		for i := 0; i < 5; i++ {
			unfoldedGroupsString = unfoldedGroupsString + groupsString
			if i < 4 {
				unfoldedGroupsString = unfoldedGroupsString + ","
			}
		}

		unfoldedGroups := utilities.ConvertStrArrToIntArr(strings.Split(unfoldedGroupsString, ","))

		result += calculatePossibilities(unfoldedPattern, unfoldedGroups)
	}
	fmt.Println(result)
}

func calculatePossibilities(pattern string, groups []int) int {

	key := pattern + " " + strings.Join(utilities.ConvertIntArrToStrArr(groups), ",")

	if cacheValue, ok := cache[key]; ok {
		return cacheValue
	}

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
	cache[key] = result
	return result
}
