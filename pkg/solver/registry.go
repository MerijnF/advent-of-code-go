package solver

import "errors"

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
