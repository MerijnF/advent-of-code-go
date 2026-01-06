package d4

import (
	"testing"

	"github.com/merijnf/advent-of-code-go/pkg/solver"
)

var INPUT = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

func TestPart1(t *testing.T) {
	ctx := solver.PuzzleContext{
		Input:    INPUT,
		Part:     1,
		Year:     2025,
		Day:      4,
		Filename: "testinput.txt",
	}

	expected := "13"
	result, err := Part1(ctx)
	if err != nil {
		t.Fatalf("Part1 failed: %v", err)
	}
	if result != expected {
		t.Errorf("Part1 returned %s, expected %s", result, expected)
	}
}

func TestPart2(t *testing.T) {
	ctx := solver.PuzzleContext{
		Input:    INPUT,
		Part:     2,
		Year:     2025,
		Day:      4,
		Filename: "testinput.txt",
	}

	expected := "43"
	result, err := Part2(ctx)
	if err != nil {
		t.Fatalf("Part2 failed: %v", err)
	}
	if result != expected {
		t.Errorf("Part2 returned %s, expected %s", result, expected)
	}
}
