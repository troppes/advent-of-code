package main

import (
	"bytes"
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
	mirror = moveStonesNorth(mirror)
	sum := 0

	for x := range mirror {
		for y := 0; y < len(mirror); y++ {
			if mirror[y][x] == 'O' {
				sum += len(mirror) - y
			}
		}
	}
	return sum
}

func part2(input string) int {

	input = strings.ReplaceAll(input, "\r\n", "\n")
	mirror := createMirror(input)

	var cache = make(map[string]int)
	end := 1000000000

	for i := 0; i < 1000000000; i++ {
		mirror = moveStonesNorth(mirror)
		mirror = moveStonesWest(mirror)
		mirror = moveStonesSouth(mirror)
		mirror = moveStonesEast(mirror)

		key := string(bytes.Join(mirror, []byte{}))
		if elem, ok := cache[key]; ok { // if found in cache
			end = i + (1000000000-i)%(i-elem) - 1
		} else { // if not add to cache
			cache[key] = i
		}

		if i == end { // if we reached the retries needed, finish
			break
		}
	}

	sum := 0

	for x := range mirror {
		for y := 0; y < len(mirror); y++ {
			if mirror[y][x] == 'O' {
				sum += len(mirror) - y
			}
		}
	}
	return sum
}

func createMirror(input string) [][]byte {

	inputArray := strings.Split(input, "\n")
	grid := make([][]byte, len(inputArray))

	for y := 0; y < len(inputArray); y++ {
		grid[y] = []byte(inputArray[y])
	}

	return grid
}

func moveStonesNorth(mirror [][]byte) [][]byte {
	for x := 0; x < len(mirror[0]); x++ {
		currCol := make([]byte, len(mirror))
		for y := 0; y < len(currCol); y++ {
			currCol[y] = mirror[y][x]
		}

		currCol = moveStones(currCol)

		for y := 0; y < len(currCol); y++ {
			mirror[y][x] = currCol[y]
		}

	}

	return mirror
}

func moveStonesWest(mirror [][]byte) [][]byte {
	for y := 0; y < len(mirror); y++ {
		mirror[y] = moveStones(mirror[y])
	}
	return mirror
}

func moveStonesEast(mirror [][]byte) [][]byte {
	for y := 0; y < len(mirror); y++ {
		currRow := make([]byte, len(mirror[y]))

		for x := 0; x < len(currRow); x++ {
			currRow[x] = mirror[y][len(currRow)-x-1]
		}
		currRow = moveStones(currRow)

		for x := 0; x < len(currRow); x++ {
			mirror[y][len(currRow)-x-1] = currRow[x]
		}
	}

	return mirror
}

func moveStonesSouth(mirror [][]byte) [][]byte {
	for x := 0; x < len(mirror[0]); x++ {
		currCol := make([]byte, len(mirror))
		for y := 0; y < len(currCol); y++ {
			currCol[y] = mirror[len(currCol)-y-1][x]
		}

		currCol = moveStones(currCol)

		for y := 0; y < len(currCol); y++ {
			mirror[len(currCol)-y-1][x] = currCol[y]
		}

	}

	return mirror
}

// instead of moving the stones differently we change the input north => south in reverse etc
func moveStones(row []byte) []byte {

	currStones := 0

	for x := len(row) - 1; x >= 0; x-- {
		if row[x] == 'O' {
			currStones++
			row[x] = '.'
		} else if row[x] == '#' {
			// cut off sand for all stones
			for i := 1; i <= currStones; i++ {
				row[x+i] = 'O'
			}
			currStones = 0
			// place new cube rock
			row[x] = '#'
		} else {
			row[x] = '.'
		}
	}
	if currStones != 0 {
		for i := 0; i < currStones; i++ {
			row[i] = 'O'
		}
	}
	return row
}
