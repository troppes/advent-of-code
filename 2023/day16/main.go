package main

import (
	"flag"
	"fmt"
	"strings"
	"sync"
	"time"

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
	start := time.Now()
	// brute force
	var wg sync.WaitGroup

	input = strings.ReplaceAll(input, "\r\n", "\n")
	dimensions, _ := createGrids(input)

	maximumChan := make(chan int)
	for x := 0; x < len(dimensions[0]); x++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			grid, hits := createGrids(input)
			maximumChan <- calculateEnergized(grid, hits, 0, x, 'D')
		}(x)

		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			grid, hits := createGrids(input)
			maximumChan <- calculateEnergized(grid, hits, len(grid)-1, x, 'U')
		}(x)
	}

	for y := 0; y < len(dimensions); y++ {
		wg.Add(1)
		go func(y int) {
			defer wg.Done()
			grid, hits := createGrids(input)
			maximumChan <- calculateEnergized(grid, hits, y, 0, 'R')
		}(y)

		wg.Add(1)
		go func(y int) {
			defer wg.Done()
			grid, hits := createGrids(input)
			maximumChan <- calculateEnergized(grid, hits, y, len(grid[0])-1, 'L')
		}(y)
	}

	go func() {
		wg.Wait()
		close(maximumChan)
	}()

	maximum := -1
	for val := range maximumChan {
		if val > maximum {
			maximum = val
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Time taken:", elapsed)
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
