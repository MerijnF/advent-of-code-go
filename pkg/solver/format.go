package solver

import (
	"fmt"
)

func FormatKey(year int, day int, part int) string {
	return fmt.Sprintf("%d-%02d-%d", year, day, part)
}

func FormatFilename(year int, day int) string {
	return fmt.Sprintf("advent_of_code_%d_%02d.txt", year, day)
}
