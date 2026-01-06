package solver

import "os"

func ReadInput(filename string) (string, error) {
	filepath := "../../input/" + filename
	input, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(input), nil
}
