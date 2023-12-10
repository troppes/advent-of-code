package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/troppes/advent-of-code/util"
)

// needed help, so used https://github.com/rumkugel13/AdventOfCode2023/blob/main/day10.go as a blueprint
func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()

	input := util.ReadInput("input_test.txt")

	if part == 1 {
		ans := part1(input)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	grid := createGrid(input)
	start := findStart(grid)

	visitedMap := map[util.Coordinate]int{start: 0}
	notChecked := []util.Coordinate{start}
	var current util.Coordinate
	maxDistance := 0

	for len(notChecked) > 0 {
		current, notChecked = util.Shift(notChecked)

		coordsToCheck := nextCoordinates(current, grid)
		for _, coord := range coordsToCheck {
			_, visited := visitedMap[coord]

			if !visited {
				visitedMap[coord] = visitedMap[current] + 1
				if visitedMap[coord] > maxDistance {
					maxDistance = visitedMap[coord]
				}
				notChecked = append(notChecked, coord)

			}
		}
	}

	return maxDistance
}

func part2(input string) int {
	grid := createGrid(input)
	start := findStart(grid)

	visitedMap := map[util.Coordinate]int{start: 0}
	notChecked := []util.Coordinate{start}
	var current util.Coordinate

	for len(notChecked) > 0 {
		current, notChecked = util.Shift(notChecked)

		coordsToCheck := nextCoordinates(current, grid)
		for _, coord := range coordsToCheck {
			_, visited := visitedMap[coord]

			if !visited {
				visitedMap[coord] = visitedMap[current] + 1
				notChecked = append(notChecked, coord)
			}
		}
	}

	countInside := 0
	for y, row := range grid {
		for x := range row {
			if isInside(grid, util.Coordinate{Y: y, X: x}, visitedMap) {
				countInside++
			}
		}
	}

	return countInside
}

func createGrid(input string) [][]string {
	inputArray := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	grid := make([][]string, len(inputArray))
	for y, line := range inputArray {
		lineArray := strings.Split(line, "")
		lineArray = append([]string{"."}, lineArray...)
		lineArray = append(lineArray, ".")

		grid[y] = make([]string, len(lineArray))
		copy(grid[y], lineArray)
	}

	padding := make([]string, len(inputArray[0])+2)
	for i := range padding {
		padding[i] = "."
	}
	grid = append([][]string{padding}, grid...)
	grid = append(grid, padding)

	return grid
}

func findStart(grid [][]string) util.Coordinate {
	for y, line := range grid {
		for x, elem := range line {
			if elem == "S" {
				return util.Coordinate{X: x, Y: y}
			}
		}
	}
	return util.Coordinate{}
}

func nextCoordinates(current util.Coordinate, grid [][]string) []util.Coordinate {

	toCheck := make([]util.Coordinate, 0)

	switch grid[current.Y][current.X] {
	case "|":
		toCheck = append(toCheck, util.Coordinate{X: current.X, Y: current.Y + 1})
		toCheck = append(toCheck, util.Coordinate{X: current.X, Y: current.Y - 1})
	case "-":
		toCheck = append(toCheck, util.Coordinate{X: current.X + 1, Y: current.Y})
		toCheck = append(toCheck, util.Coordinate{X: current.X - 1, Y: current.Y})
	case "L":
		toCheck = append(toCheck, util.Coordinate{X: current.X, Y: current.Y - 1})
		toCheck = append(toCheck, util.Coordinate{X: current.X + 1, Y: current.Y})
	case "J":
		toCheck = append(toCheck, util.Coordinate{X: current.X, Y: current.Y - 1})
		toCheck = append(toCheck, util.Coordinate{X: current.X - 1, Y: current.Y})
	case "7":
		toCheck = append(toCheck, util.Coordinate{X: current.X, Y: current.Y + 1})
		toCheck = append(toCheck, util.Coordinate{X: current.X - 1, Y: current.Y})
	case "F":
		toCheck = append(toCheck, util.Coordinate{X: current.X, Y: current.Y + 1})
		toCheck = append(toCheck, util.Coordinate{X: current.X + 1, Y: current.Y})
	case "S":
		north := grid[current.Y+1][current.X]
		south := grid[current.Y-1][current.X]
		west := grid[current.Y][current.X+1]
		east := grid[current.Y][current.X-1]

		if north == "|" || north == "L" || north == "J" {
			toCheck = append(toCheck, util.Coordinate{Y: current.Y + 1, X: current.X})
		}
		if south == "|" || south == "7" || south == "F" {
			toCheck = append(toCheck, util.Coordinate{Y: current.Y - 1, X: current.X})
		}
		if west == "-" || west == "J" || west == "7" {
			toCheck = append(toCheck, util.Coordinate{Y: current.Y, X: current.X + 1})
		}
		if east == "-" || east == "L" || east == "F" {
			toCheck = append(toCheck, util.Coordinate{Y: current.Y, X: current.X - 1})
		}
	}

	return toCheck
}

func isInside(grid [][]string, coord util.Coordinate, theLoop map[util.Coordinate]int) bool {
	// check if its parts of the loop
	if _, part := theLoop[coord]; part {
		return false
	}

	count := 0
	cornerCounts := map[string]int{}
	for y := coord.Y + 1; y < len(grid); y++ { // only look up, since the parity does not need more information
		check := util.Coordinate{X: coord.X, Y: y}
		tile := grid[y][coord.X]

		// determine the type of the starting point
		if tile == "S" {
			tile = findStartType(util.Coordinate{X: coord.X, Y: y}, grid)
		}

		if _, part := theLoop[check]; part {
			if tile == "-" { // we scan if we have a ceiling: -
				count++
			} else if tile != "|" && tile != "." { // if not and we have a piece that connect coners we add it
				cornerCounts[tile]++
			}
		}
	}

	// add the surplus of corners to the count
	count += max(cornerCounts["L"], cornerCounts["7"]) - util.Abs(cornerCounts["L"]-cornerCounts["7"])
	count += max(cornerCounts["F"], cornerCounts["J"]) - util.Abs(cornerCounts["F"]-cornerCounts["J"])

	// check parity if we are inside or outside, if even, then we are outside
	return count%2 == 1
}

func findStartType(start util.Coordinate, grid [][]string) string {
	points := nextCoordinates(start, grid) // determine surroundings

	minX := min(points[0].X, points[1].X)
	maxX := max(points[0].X, points[1].X)
	minY := min(points[0].Y, points[1].Y)
	maxY := max(points[0].Y, points[1].Y)

	if points[0].X == points[1].X { // if both x are in a line, it can only go up
		return "|"
	} else if points[0].Y == points[1].Y { // if both y are in a line, it can only go sideways
		return "-"
	} else if minX < start.X && minY < start.Y { // if both are smaller we move down and left
		return "J"
	} else if maxX > start.X && maxY > start.Y { // if both are higher we move one up and right
		return "F"
	} else if maxX > start.X && minY < start.Y { // if Y is smaller we go down, if X is higher we move right
		return "L"
	} else if minX < start.X && maxY > start.Y { // if Y is higher we go up, if X is smaller we move left
		return "7"
	}
	return "." // cannot happen (normally)
}
