package main

import (
	"fmt"

	"github.com/merijnf/advent-of-code-go/utilities"
)

type instruction rune
type node struct {
	key   string
	left  string
	right string
}
type phaseAndPeriod struct {
	phase  int
	period int
}

const start = "A"
const end = "Z"

const (
	L instruction = 'L'
	R instruction = 'R'
)

func main() {
	lines := utilities.LoadInputStringLines(2023, 8)

	instructions := parseInstructions(lines[0])
	nodesLines := lines[2:]
	nodes := make(map[string]node, len(nodesLines))
	for _, nodeLine := range nodesLines {
		node := parseNode(nodeLine)
		nodes[node.key] = node
	}

	routes := make([]string, 0)
	for k := range nodes {
		if k[2:] == start {
			routes = append(routes, k)
		}
	}

	frequencies := make([]phaseAndPeriod, len(routes))

	fmt.Println(routes)

	for index, route := range routes {
		currentNode := route
		steps := 0
		endsFirstSeenAt := make(map[string]int, 0)
		for i := 0; i < len(instructions); i++ {
			steps++
			switch instructions[i] {
			case R:
				currentNode = nodes[currentNode].right
			case L:
				currentNode = nodes[currentNode].left
			}
			if currentNode[2:] == end {
				val, ok := endsFirstSeenAt[currentNode]
				if ok {
					frequencies[index] = phaseAndPeriod{phase: val, period: steps - val}
					break
				} else {
					endsFirstSeenAt[currentNode] = steps
				}
			}
			if i == len(instructions)-1 {
				i = -1
			}

		}
	}

	fmt.Println(frequencies)

	//lcm, will not work for all inputs
	phases := make([]int, len(frequencies))
	for i, frequency := range frequencies {
		phases[i] = frequency.phase
	}
	value := lcm(phases[0], phases[1], phases[2:]...)
	fmt.Println(value)

	// BRUTE FORCE
	// steps := 0
	// for i := 0; i < len(instructions); i++ {
	// 	steps++
	// 	for j, route := range routes {
	// 		switch instructions[i] {
	// 		case R:
	// 			routes[j] = nodes[route].right
	// 		case L:
	// 			routes[j] = nodes[route].left
	// 		}
	// 	}
	// 	if allRoutesEnd(routes) {
	// 		break
	// 	}

	// 	if i == len(instructions)-1 {
	// 		i = -1
	// 	}
	// }
	// fmt.Println(steps)
}

func parseInstructions(line string) []instruction {
	return []instruction(line)
}

func parseNode(line string) node {
	key := line[0:3]
	left := line[7:10]
	right := line[12:15]
	return node{key: key, left: left, right: right}
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}
	return result
}

// func allRoutesEnd(routes []string) bool {
// 	for _, route := range routes {
// 		if route[2:] != end {
// 			return false
// 		}
// 	}
// 	return true
// }
