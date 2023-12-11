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
	grid := createAndExpandGrid(input)

	galaxies := findGalaxies(grid, 2)
	numDistances := len(galaxies) * (len(galaxies) - 1) / 2

	distances := make([]int, numDistances)
	idx := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			distances[idx] = galaxies[i].ManhattanDistance(&galaxies[j])
			idx++
		}
	}

	return util.Reduce(distances, func(a int, b int) int {
		return a + b
	}, 0)
}

func part2(input string) int {
	grid := createAndExpandGrid(input)

	galaxies := findGalaxies(grid, 1000000)
	numDistances := len(galaxies) * (len(galaxies) - 1) / 2

	distances := make([]int, numDistances)
	idx := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			distances[idx] = galaxies[i].ManhattanDistance(&galaxies[j])
			idx++
		}
	}

	return util.Reduce(distances, func(a int, b int) int {
		return a + b
	}, 0)
}

func createAndExpandGrid(input string) [][]string {
	inputArray := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	grid := make([][]string, 0)

	for y := 0; y < len(inputArray); y++ {
		lineArray := strings.Split(inputArray[y], "")

		empty := true
		for x := 0; x < len(lineArray); x++ {
			// if only ., dann add new line before
			if lineArray[x] == "#" {
				empty = false
				break
			}
		}

		// if emtpy then add padding
		if empty {
			emptyLine := make([]string, len(lineArray))
			for x := 0; x < len(lineArray); x++ {
				emptyLine[x] = "X"
			}
			grid = append(grid, emptyLine)
		}
		grid = append(grid, lineArray)
	}

	// new grid to store updated values
	newGrid := make([][]string, len(grid))
	for i := range grid {
		newGrid[i] = make([]string, len(grid[0]))
		copy(newGrid[i], grid[i])
	}

	offset := 0
	for x := 0; x < len(grid[0]); x++ {

		empty := true
		for y := 0; y < len(grid); y++ {
			if grid[y][x] == "#" {
				empty = false
				break
			}
		}

		if empty {
			// Append the slice to the specific column
			for i := range newGrid {
				pre := newGrid[i][:x+1+offset]
				after := newGrid[i][x+1+offset:]

				newGrid[i] = append(pre, append([]string{"X"}, after...)...)
			}
			offset++
		}
	}

	return newGrid
}

func findGalaxies(grid [][]string, expansion int) []util.Coordinate {
	var galaxies []util.Coordinate

	for y, line := range grid {
		for x, elem := range line {
			if elem == "#" {
				galaxies = append(galaxies, util.Coordinate{X: x, Y: y})
			}
		}
	}

	for i, galaxy := range galaxies {
		newX, newY := 0, 0

		for x := 0; x < galaxy.X; x++ {
			if grid[galaxy.Y][x] == "X" {
				newX += expansion - 2
			}
			newX++
		}

		for y := 0; y < galaxy.Y; y++ {
			if grid[y][galaxy.X] == "X" {
				newY += expansion - 2
			}
			newY++
		}

		galaxies[i] = util.Coordinate{X: newX, Y: newY}
	}

	return galaxies
}
