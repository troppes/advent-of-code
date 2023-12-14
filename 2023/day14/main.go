package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/troppes/advent-of-code/util"
)

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
	mirror := createMirror(input)
	mirror = calcStonePositions(mirror)
	sum := 0

	for _, col := range mirror {
		for i := len(col) - 1; i >= 0; i-- {
			if col[i] == 'O' {
				sum += i + 1
			}
		}
	}

	return sum
}

func part2(input string) int {
	return part1(input)
}

func createMirror(input string) []string {
	rows := strings.Split(input, "\n")
	cols := make([]string, len(rows[0]))

	for x := 0; x < len(rows[0]); x++ {
		var sb strings.Builder
		for y := 0; y < len(rows); y++ {
			sb.WriteByte(rows[y][x])
		}
		cols[x] = sb.String()
	}

	return cols
}

func calcStonePositions(mirror []string) []string {

	newMirrors := make([]string, len(mirror))

	// if the same stop otherwise
	for y := 0; y < len(mirror); y++ {
		var sb strings.Builder

		currStones := 0
		currCol := mirror[y]
		for x := len(currCol) - 1; x >= 0; x-- { // change direction to change between south and north
			if currCol[x] == 'O' {
				currStones++
				sb.WriteByte('.')
			} else if currCol[x] == '#' {
				// cut off sand for all stones
				currString := sb.String()
				sb.Reset()
				sb.WriteString(currString[:len(currString)-currStones])
				for i := 0; i < currStones; i++ {
					sb.WriteByte('O')
				}
				currStones = 0
				// place new cube rock
				sb.WriteByte('#')
			} else {
				sb.WriteByte('.')
			}
		}
		if currStones != 0 {
			currString := sb.String()
			sb.Reset()
			sb.WriteString(currString[:len(currString)-currStones])
			for i := 0; i < currStones; i++ {
				sb.WriteByte('O')
			}
		}

		newMirrors[y] = sb.String()
	}

	return newMirrors
}
