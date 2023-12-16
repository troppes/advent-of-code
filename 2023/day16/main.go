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
	// just to be sure remove all newlines
	input = strings.ReplaceAll(input, "\r\n", "\n")
	grid, hits := createGrids(input)

	return calculateEnergized(grid, hits, 0, 0, 'R')
}

func part2(input string) int {
	// brute force
	input = strings.ReplaceAll(input, "\r\n", "\n")
	dimensions, _ := createGrids(input)

	maximum := -1
	for x := 0; x < len(dimensions[0]); x++ {
		grid, hits := createGrids(input)
		c1 := calculateEnergized(grid, hits, 0, x, 'D')

		grid, hits = createGrids(input)
		c2 := calculateEnergized(grid, hits, len(grid)-1, x, 'U')

		currMax := max(c1, c2)
		if currMax > maximum {
			maximum = currMax
		}
	}

	for y := 0; y < len(dimensions); y++ {
		grid, hits := createGrids(input)
		c1 := calculateEnergized(grid, hits, y, 0, 'R')

		grid, hits = createGrids(input)
		c2 := calculateEnergized(grid, hits, y, len(grid[0])-1, 'L')

		currMax := max(c1, c2)
		if currMax > maximum {
			maximum = currMax
		}
	}

	return maximum

}

func calculateEnergized(grid [][]byte, hits [][]bool, sY, sX int, startingDirection rune) int {
	visited := make(map[string]bool)

	determineHits(&hits, &visited, grid, sY, sX, startingDirection)

	count := 0
	for _, row := range hits {
		for _, val := range row {
			if val {
				count++
			}
		}
	}
	return count

}

func createGrids(input string) ([][]byte, [][]bool) {

	inputArray := strings.Split(input, "\n")
	grid := make([][]byte, len(inputArray))
	boolGrid := make([][]bool, len(inputArray))

	for y := 0; y < len(inputArray); y++ {
		grid[y] = []byte(inputArray[y])
		boolGrid[y] = make([]bool, len(inputArray[y]))
	}

	return grid, boolGrid
}

func determineHits(hits *[][]bool, visited *map[string]bool, input [][]byte, y, x int, currDirection rune) {
	key := fmt.Sprintf("y%d_x%d_c%c", y, x, currDirection)

	if y < 0 || y >= len(input) || x < 0 || x >= len(input[y]) { // check oob
		return
	}

	if (*visited)[key] { // we may loop indefinitely, so we need to check if we were at this position with this direction
		return
	}

	(*visited)[key] = true
	(*hits)[y][x] = true

	if currDirection == 'R' { // right
		if input[y][x] == '/' {
			determineHits(hits, visited, input, y-1, x, 'U') // up
		} else if input[y][x] == '\\' {
			determineHits(hits, visited, input, y+1, x, 'D') // down
		} else if input[y][x] == '|' {
			determineHits(hits, visited, input, y-1, x, 'U') // up
			determineHits(hits, visited, input, y+1, x, 'D') // down
		} else {
			determineHits(hits, visited, input, y, x+1, 'R') // right
		}
	} else if currDirection == 'L' { // left
		if input[y][x] == '/' {
			determineHits(hits, visited, input, y+1, x, 'D') // down
		} else if input[y][x] == '\\' {
			determineHits(hits, visited, input, y-1, x, 'U') // up
		} else if input[y][x] == '|' {
			determineHits(hits, visited, input, y-1, x, 'U') // up
			determineHits(hits, visited, input, y+1, x, 'D') // down
		} else {
			determineHits(hits, visited, input, y, x-1, 'L') // left
		}
	} else if currDirection == 'U' { // up
		if input[y][x] == '/' {
			determineHits(hits, visited, input, y, x+1, 'R') // right
		} else if input[y][x] == '\\' {
			determineHits(hits, visited, input, y, x-1, 'L') // left

		} else if input[y][x] == '-' {
			determineHits(hits, visited, input, y, x-1, 'L') // left
			determineHits(hits, visited, input, y, x+1, 'R') // right
		} else {
			determineHits(hits, visited, input, y-1, x, 'U') // up
		}
	} else if currDirection == 'D' { // down
		if input[y][x] == '/' {
			determineHits(hits, visited, input, y, x-1, 'L') // left
		} else if input[y][x] == '\\' {
			determineHits(hits, visited, input, y, x+1, 'R') // right
		} else if input[y][x] == '-' {
			determineHits(hits, visited, input, y, x-1, 'L') // left
			determineHits(hits, visited, input, y, x+1, 'R') // right
		} else {
			determineHits(hits, visited, input, y+1, x, 'D') // down
		}
	}
}
