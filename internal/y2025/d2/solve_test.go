package d2

import (
	"testing"

	"github.com/merijnf/advent-of-code-go/pkg/solver"
)

var INPUT = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

func TestPart1(t *testing.T) {
	ctx := solver.PuzzleContext{
		Input:    INPUT,
		Part:     1,
		Year:     2025,
		Day:      2,
		Filename: "testinput.txt",
	}

	expected := "1227775554"
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
		Day:      2,
		Filename: "testinput.txt",
	}

	expected := "4174379265"
	result, err := Part2(ctx)
	if err != nil {
		t.Fatalf("Part2 failed: %v", err)
	}
	if result != expected {
		t.Errorf("Part2 returned %s, expected %s", result, expected)
	}
}
