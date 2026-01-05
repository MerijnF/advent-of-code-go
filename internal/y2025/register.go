package y2025

import (
	"github.com/merijnf/advent-of-code-go/internal/y2025/d1"
	"github.com/merijnf/advent-of-code-go/internal/y2025/d2"
	"github.com/merijnf/advent-of-code-go/pkg/solver"
)

func RegisterSolvers() {
	// Register day 1 solvers
	solver.RegisterSolver(solver.FormatKey(2025, 1, 1), d1.Part1)
	solver.RegisterSolver(solver.FormatKey(2025, 1, 2), d1.Part2)
	// Register day 2 solvers
	solver.RegisterSolver(solver.FormatKey(2025, 2, 1), d2.Part1)
	solver.RegisterSolver(solver.FormatKey(2025, 2, 2), d2.Part2)
}
