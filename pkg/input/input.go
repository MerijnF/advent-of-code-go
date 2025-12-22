package input

import (
	"fmt"
	"os"
	"strconv"
)

func ReadInput(year int, day int, testMode bool) string {
	filename := "../../input/advent_of_code_" + strconv.Itoa(year) + "_" + fmt.Sprintf("%02d", day)
	if testMode {
		filename = filename + "_test"
	}
	filename = filename + ".txt"
	input, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(input)
}
