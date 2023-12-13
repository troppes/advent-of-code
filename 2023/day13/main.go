package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/troppes/advent-of-code/util"
)

type Mirror struct {
	cols []string
	rows []string
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
	sum := 0

	for _, input := range inputArray {
		mirror := createMirror(input)
		sum += findReflections(mirror.rows, 100)
		sum += findReflections(mirror.cols, 1)
	}

	return sum
}

func part2(input string) int {
	inputArray := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n\n")
	sum := 0

	for _, input := range inputArray {
		mirror := createMirror(input)
		sum += findReflectionsAndFixSmudge(mirror.rows, 100)
		sum += findReflectionsAndFixSmudge(mirror.cols, 1)
	}

	return sum
}

func createMirror(input string) Mirror {
	rows := strings.Split(input, "\n")
	cols := make([]string, len(rows[0]))

	for x := 0; x < len(rows[0]); x++ {
		var sb strings.Builder
		for y := 0; y < len(rows); y++ {
			sb.WriteByte(rows[y][x])
		}
		cols[x] += sb.String()
	}

	return Mirror{rows: rows, cols: cols}
}

func findReflections(lines []string, multiplier int) int {
	for i := 1; i < len(lines); i++ {
		if lines[i] != lines[i-1] { // if its not the same skip
			continue
		}
		// otherwise search
		x, y := i, i-1
		isValid := true

		for j := 0; j < min(i, len(lines)-i)-1; j++ { // split the indices and go seraching
			x++
			y--
			if lines[x] != lines[y] {
				isValid = false
				break
			}
		}

		if isValid {
			return i * multiplier
		}
	}

	return 0
}

func findReflectionsAndFixSmudge(lines []string, multiplier int) int {
	for i := 1; i < len(lines); i++ {
		differences := differencesInLine(lines[i], lines[i-1])
		if differences > 1 { // if we have more than one smudge stop
			continue
		}

		x, y := i, i-1
		isValid := true

		for j := 0; j < min(i, len(lines)-i)-1; j++ {
			x++
			y--
			differences += differencesInLine(lines[x], lines[y])
			if differences > 1 {
				isValid = false
				break
			}
		}

		if isValid && differences == 1 { // we need to check if we fixed a smudge, otherwise its not valid
			return i * multiplier
		}
	}

	return 0
}

func differencesInLine(s1, s2 string) int {
	if len(s1) != len(s2) {
		panic("Strings are not same size!")
	}
	count := 0
	for i := range s1 {
		if s1[i] != s2[i] {
			count++
		}
	}
	return count
}
