package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/troppes/advent-of-code/util"
)

type Line struct {
	p1     util.Coordinate
	p2     util.Coordinate
	length int
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
	input = strings.ReplaceAll(input, "\r\n", "\n")
	lines := convertInputToLines(input)
	result := 0

	// https://en.wikipedia.org/wiki/Shoelace_formula
	for i := 0; i < len(lines); i++ {
		result += (lines[i].p1.Y+lines[i].p2.Y)*(lines[i].p1.X-lines[i].p2.X) + lines[i].length
	}

	return result/2 + 1
}

func part2(input string) int {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	lines := convertHexInputToLines(input)
	result := 0

	for i := 0; i < len(lines); i++ {
		result += (lines[i].p1.Y+lines[i].p2.Y)*(lines[i].p1.X-lines[i].p2.X) + lines[i].length
	}

	return result/2 + 1
}

func convertInputToLines(input string) []Line {
	inputArray := strings.Split(input, "\n")

	lines := make([]Line, len(inputArray))

	x1, x2, y1, y2 := 0, 0, 0, 0

	for i, line := range inputArray {
		parts := strings.Split(line, " ")

		num, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		switch parts[0] {
		case "R":
			x2 = x1 + num
		case "L":
			x2 = x1 - num
		case "D":
			y2 = y1 + num
		case "U":
			y2 = y1 - num
		}

		p1 := util.Coordinate{X: x1, Y: y1}
		p2 := util.Coordinate{X: x2, Y: y2}
		distance := p1.ManhattanDistance(&p2)

		lines[i] = Line{p1: p1, p2: p2, length: distance}

		x1, y1 = x2, y2 // make the new startin point the old end point
	}
	return lines
}

func convertHexInputToLines(input string) []Line {
	inputArray := strings.Split(input, "\n")

	lines := make([]Line, len(inputArray))

	x1, x2, y1, y2 := 0, 0, 0, 0

	for i, line := range inputArray {
		parts := strings.Split(line, " ")

		hexString := parts[2]
		hexNum := hexString[2:7]
		direction := hexString[7]

		num, err := strconv.ParseInt(hexNum, 16, 32) // 32 still return 64 int
		if err != nil {
			panic(err)
		}
		num32 := int(num) // needed since currently my coordinates are int32

		switch direction {
		case '0':
			x2 = x1 + num32
		case '1':
			y2 = y1 + num32
		case '2':
			x2 = x1 - num32
		case '3':
			y2 = y1 - num32
		}

		p1 := util.Coordinate{X: x1, Y: y1}
		p2 := util.Coordinate{X: x2, Y: y2}
		distance := p1.ManhattanDistance(&p2)

		lines[i] = Line{p1: p1, p2: p2, length: distance}

		x1, y1 = x2, y2 // make the new startin point the old end point
	}
	return lines
}
