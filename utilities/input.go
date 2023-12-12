package utilities

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LoadInputString(year int, day int) string {
	yearStr := strconv.Itoa(year)
	dayStr := strconv.Itoa(day)
	if len(dayStr) == 1 {
		dayStr = "0" + dayStr
	}

	fmt.Println("loading input data " + yearStr + " day" + dayStr)
	//dat, err := os.ReadFile("input-test.txt")
	dat, err := os.ReadFile("../../input/" + yearStr + "/day" + dayStr + "/input.txt")
	CheckError(err)
	return string(dat)
}

func LoadInputStringLines(year int, day int) []string {
	input := LoadInputString(year, day)
	return strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
}
