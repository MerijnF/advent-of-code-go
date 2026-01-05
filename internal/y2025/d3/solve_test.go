package d3

import (
	"testing"

	"github.com/merijnf/advent-of-code-go/pkg/solver"
)

func TestPart2(t *testing.T) {
	ctx := solver.PuzzleContext{
		Input:    "987654321111111\r\n811111111111119\r\n234234234234278\r\n818181911112111",
		Part:     2,
		Year:     2025,
		Day:      3,
		Filename: "testinput.txt",
	}

	expected := "3121910778619"
	result, err := Part2(ctx)
	if err != nil {
		t.Fatalf("Part2 failed: %v", err)
	}
	if result != expected {
		t.Errorf("Part2 returned %s, expected %s", result, expected)
	}
}
