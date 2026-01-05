package solver

import (
	"errors"
	"fmt"
	"os"
)

type PuzzleContext struct {
	Key      string
	Year     int
	Day      int
	Part     int
	Input    string
	Filename string
}

func Solve(year int, day int, part int) (string, error) {
	key := FormatKey(year, day, part)
	solver, err := GetSolver(key)
	if err != nil {
		return "no solver", err
	}
	filename := FormatFilename(year, day)
	input, err := ReadInput(filename)
	if err != nil {
		return "", err
	}
	solverCtx := PuzzleContext{
		Key:      key,
		Filename: filename,
		Year:     year,
		Day:      day,
		Part:     part,
		Input:    input,
	}
	return solver(solverCtx)
}

var solvers = make(map[string]func(ctx PuzzleContext) (string, error))

func GetSolver(key string) (func(ctx PuzzleContext) (string, error), error) {
	solver, exists := solvers[key]
	if !exists {
		return nil, errors.New("no solver registered for key: " + key)
	}
	return solver, nil
}

func RegisterSolver(key string, solver func(ctx PuzzleContext) (string, error)) {
	solvers[key] = solver
}

func FormatKey(year int, day int, part int) string {
	return fmt.Sprintf("%d-%02d-%d", year, day, part)
}

func FormatFilename(year int, day int) string {
	return fmt.Sprintf("advent_of_code_%d_%02d.txt", year, day)
}

func ReadInput(filename string) (string, error) {
	filepath := "../../input/" + filename
	input, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(input), nil
}
