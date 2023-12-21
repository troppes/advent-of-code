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
	garden := createGarden(input)
	startingPoint := findStart(garden)

	return takeAWalk(garden, startingPoint, 64)
}

func part2(input string) int {
	return part1(input)
}

func findStart(garden [][]byte) util.Coordinate {
	for y := 0; y < len(garden); y++ {
		for x := 0; x < len(garden[y]); x++ {
			if garden[y][x] == 'S' {
				return util.Coordinate{X: x, Y: y}
			}
		}
	}
	return util.Coordinate{X: -1, Y: -1}
}

func takeAWalk(garden [][]byte, start util.Coordinate, maxMoves int) int {
	var visited = make(map[int][]util.Coordinate)
	visited[0] = append(visited[0], start)

	directions := []util.Direction{util.Up, util.Down, util.Left, util.Right}

	for move := 0; move < maxMoves; move++ {
		for _, currentCoord := range visited[move] {
			for _, dir := range directions {
				newCoord := currentCoord.Move(dir, 1)
				if isInBounds(garden, newCoord) && !util.Contains(newCoord, visited[move+1]) {
					visited[move+1] = append(visited[move+1], newCoord)
				}
			}
		}
	}

	sum := len(visited[len(visited)-1])
	return sum
}

func isInBounds(garden [][]byte, coord util.Coordinate) bool {
	if coord.X < 0 || coord.X > len(garden[0])-1 || coord.Y < 0 || coord.Y > len(garden)-1 || garden[coord.Y][coord.X] == '#' {
		return false
	}
	return true
}

func createGarden(input string) [][]byte {

	inputArray := strings.Split(input, "\n")
	grid := make([][]byte, len(inputArray))

	for y := 0; y < len(inputArray); y++ {
		grid[y] = []byte(inputArray[y])
	}

	return grid
}
