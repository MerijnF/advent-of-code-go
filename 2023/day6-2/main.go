package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/merijnf/advent-of-code-go/utilities"
)

func main() {
	lines := utilities.LoadInputStringLines(2023, 6)

	time, error := strconv.Atoi(strings.ReplaceAll(strings.Split(lines[0], ":")[1], " ", ""))
	utilities.CheckError(error)
	distance, error := strconv.Atoi(strings.ReplaceAll(strings.Split(lines[1], ":")[1], " ", ""))
	utilities.CheckError(error)

	// d: distance traveled
	// t: time
	// x: held time
	// d = x * (t - x)
	// d = tx - x^2
	// x^2 - tx + d = 0
	// quadratic formula
	// x = t+(sqrt(t^2-4d))/2
	// x = t-(sqrt(t^2-4d))/2

	d := float64(distance)
	t := float64(time)

	fmt.Println("solving for d: " + strconv.Itoa(distance) + " t: " + strconv.Itoa(time))

	x1 := int(math.Ceil(((t + math.Sqrt(math.Pow(t, 2)-4*d)) / 2) - 1))
	x2 := int(math.Floor(((t - math.Sqrt(math.Pow(t, 2)-4*d)) / 2) + 1))
	fmt.Println(x1)
	fmt.Println(x2)
	solves := x1 - x2 + 1

	fmt.Println("final value: " + strconv.Itoa(solves))
}
