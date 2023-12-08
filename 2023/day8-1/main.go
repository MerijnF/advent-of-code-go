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

const start = "AAA"
const end = "ZZZ"

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
	fmt.Println(instructions)
	fmt.Println(nodes)

	steps := 0
	currentNode := start
	for i := 0; i < len(instructions); i++ {
		steps++
		switch instructions[i] {
		case R:
			currentNode = nodes[currentNode].right
		case L:
			currentNode = nodes[currentNode].left
		}
		if currentNode == end {
			break
		}
		if i == len(instructions)-1 {
			i = -1
		}
	}
	fmt.Println(steps)
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
