package main

import (
	"flag"
	"fmt"
	"regexp"
	"strings"

	"github.com/troppes/advent-of-code/util"
)

type TreeNode struct {
	value string
	left  string
	right string
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()

	input := util.ReadInput("input.txt")

	if part == 1 {
		ans := part1(input)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {

	inputArray := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n\n")

	instructions := strings.Split(inputArray[0], "")
	dataStructure := strings.Split(inputArray[1], "\n")

	treeMap, _ := createTree(dataStructure)

	current := treeMap["AAA"]
	steps := 0
	i := 0
	for current.value != "ZZZ" {
		if instructions[i] == "L" {
			current = treeMap[current.left]
		} else {
			current = treeMap[current.right]
		}
		i++
		steps++
		if i >= len(instructions) {
			i = 0
		}
	}

	return steps
}

func part2(input string) int {
	inputArray := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n\n")

	instructions := strings.Split(inputArray[0], "")
	dataStructure := strings.Split(inputArray[1], "\n")

	treeMap, currentNodes := createTree(dataStructure)

	results := make([]int, len(currentNodes))

	for j := 0; j < len(currentNodes); j++ {
		current := currentNodes[j]
		steps := 0
		i := 0
		for !strings.HasSuffix(current.value, "Z") {
			if instructions[i] == "L" {
				current = treeMap[current.left]
			} else {
				current = treeMap[current.right]
			}
			i++
			steps++
			if i >= len(instructions) {
				i = 0
			}
		}
		results[j] = steps
	}

	solution := results[0]

	for i := 1; i < len(results); i++ {
		solution = lcm(solution, results[i])
	}

	return solution
}

func lcm(a, b int) int {
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}(a, b)

	return (a * b) / gcd
}

func createTree(input []string) (map[string]*TreeNode, []*TreeNode) {
	treeMap := make(map[string]*TreeNode)
	startingNodes := make([]*TreeNode, 0)

	for _, line := range input {
		pattern := regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)
		matches := pattern.FindStringSubmatch(line)

		node := new(TreeNode)
		node.value = matches[1]
		node.left = matches[2]
		node.right = matches[3]
		treeMap[node.value] = node

		if strings.HasSuffix(node.value, "A") {
			startingNodes = append(startingNodes, node)
		}
	}

	return treeMap, startingNodes
}
