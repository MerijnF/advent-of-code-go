package d4

import (
	"strconv"

	"github.com/merijnf/advent-of-code-go/pkg/solver"
	"github.com/merijnf/advent-of-code-go/pkg/utils/strutil"
)

const empty = '.'
const filled = '@'

func Part1(ctx solver.PuzzleContext) (string, error) {
	lines := strutil.SplitLinesAndNormalize(ctx.Input, false)

	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	accessibleCount := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == filled {

				filledAdjecent := 0
				// check up
				if y > 0 && grid[y-1][x] == filled {
					filledAdjecent++
				}
				// check down
				if y < len(grid)-1 && grid[y+1][x] == filled {
					filledAdjecent++
				}
				// check left
				if x > 0 && grid[y][x-1] == filled {
					filledAdjecent++
				}
				// check right
				if x < len(grid[y])-1 && grid[y][x+1] == filled {
					filledAdjecent++
				}
				// check up-left
				if y > 0 && x > 0 && grid[y-1][x-1] == filled {
					filledAdjecent++
				}
				// check up-right
				if y > 0 && x < len(grid[y])-1 && grid[y-1][x+1] == filled {
					filledAdjecent++
				}
				// check down-left
				if y < len(grid)-1 && x > 0 && grid[y+1][x-1] == filled {
					filledAdjecent++
				}
				// check down-right
				if y < len(grid)-1 && x < len(grid[y])-1 && grid[y+1][x+1] == filled {
					filledAdjecent++
				}
				if filledAdjecent < 4 {
					accessibleCount++
				}
			}
		}
	}

	return strconv.Itoa(accessibleCount), nil
}

func Part2(ctx solver.PuzzleContext) (string, error) {
	lines := strutil.SplitLinesAndNormalize(ctx.Input, false)

	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	removedCount := 0
	accessibleCount := 1 // dummy to enter the loop
	for accessibleCount > 0 {
		accessibleCount = 0
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				if grid[y][x] == filled {

					filledAdjecent := 0
					// check up
					if y > 0 && grid[y-1][x] == filled {
						filledAdjecent++
					}
					// check down
					if y < len(grid)-1 && grid[y+1][x] == filled {
						filledAdjecent++
					}
					// check left
					if x > 0 && grid[y][x-1] == filled {
						filledAdjecent++
					}
					// check right
					if x < len(grid[y])-1 && grid[y][x+1] == filled {
						filledAdjecent++
					}
					// check up-left
					if y > 0 && x > 0 && grid[y-1][x-1] == filled {
						filledAdjecent++
					}
					// check up-right
					if y > 0 && x < len(grid[y])-1 && grid[y-1][x+1] == filled {
						filledAdjecent++
					}
					// check down-left
					if y < len(grid)-1 && x > 0 && grid[y+1][x-1] == filled {
						filledAdjecent++
					}
					// check down-right
					if y < len(grid)-1 && x < len(grid[y])-1 && grid[y+1][x+1] == filled {
						filledAdjecent++
					}
					if filledAdjecent < 4 {
						accessibleCount++
						grid[y][x] = empty
						removedCount++
					}
				}
			}
		}
	}

	return strconv.Itoa(removedCount), nil
}
