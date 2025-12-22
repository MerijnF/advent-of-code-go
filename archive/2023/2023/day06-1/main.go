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

	times := utilities.ConvertStrArrToIntArr(strings.Fields(strings.Split(lines[0], ":")[1]))
	distances := utilities.ConvertStrArrToIntArr(strings.Fields(strings.Split(lines[1], ":")[1]))

	fmt.Println(times)
	fmt.Println(distances)

	// d: distance traveled
	// t: time
	// x: held time
	// d = x * (t - x)
	// d = tx - x^2
	// x^2 - tx + d = 0
	// quadratic formula
	// x = t+(sqrt(t^2-4d))/2
	// x = t-(sqrt(t^2-4d))/2
	result := 1
	for i := 0; i < len(times); i++ {

		d := float64(distances[i])
		t := float64(times[i])

		fmt.Println("solving for d: " + strconv.Itoa(distances[i]) + " t: " + strconv.Itoa(times[i]))

		x1 := int(math.Ceil(((t + math.Sqrt(math.Pow(t, 2)-4*d)) / 2) - 1))
		x2 := int(math.Floor(((t - math.Sqrt(math.Pow(t, 2)-4*d)) / 2) + 1))
		fmt.Println(x1)
		fmt.Println(x2)
		solves := x1 - x2 + 1
		result = result * solves
	}
	fmt.Println("final value: " + strconv.Itoa(result))
}
