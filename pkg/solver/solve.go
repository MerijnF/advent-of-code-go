package solver

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
