package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/oleiade/lane/v2" // prio queue
	"github.com/troppes/advent-of-code/util"
)

type QueueEntry struct {
	pos      util.Coordinate
	dir      util.Direction
	straight int
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
	grid, finalCord := parseGrid(input)
	return findBestPath(grid, finalCord, 3, 0)
}

func part2(input string) int {
	grid, finalCord := parseGrid(input)
	return findBestPath(grid, finalCord, 10, 4)
}

func parseGrid(input string) (map[util.Coordinate]int, util.Coordinate) {
	inputArray := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	coordinates := make(map[util.Coordinate]int)
	finalCord := util.Coordinate{X: -1, Y: -1}

	for y, line := range inputArray {
		for x, character := range line {
			num, err := strconv.Atoi(string(character))
			if err != nil {
				panic(err)
			}
			finalCord = util.Coordinate{X: x, Y: y}
			coordinates[util.Coordinate{X: x, Y: y}] = num
		}
	}

	return coordinates, finalCord
}

// BFS to find the best path modified with an prioq
func findBestPath(board map[util.Coordinate]int, target util.Coordinate, maxStraight int, minStraight int) int {

	priorityQueue := lane.NewMinPriorityQueue[QueueEntry, int]()

	priorityQueue.Push(QueueEntry{
		pos:      util.Coordinate{Y: 0, X: 1},
		straight: 1,
		dir:      util.Right,
	}, 0)
	priorityQueue.Push(QueueEntry{
		pos:      util.Coordinate{Y: 1, X: 0},
		straight: 1,
		dir:      util.Down,
	}, 0)
	cache := make(map[QueueEntry]int)

	for !priorityQueue.Empty() {
		currElem, heatLoss, _ := priorityQueue.Pop()

		if _, exists := board[currElem.pos]; !exists { // oob check
			continue
		}

		heatLoss += board[currElem.pos]
		if currElem.pos == target {
			return heatLoss
		}

		if v, exists := cache[currElem]; exists { // need cache otherwise running to long
			if v <= heatLoss { // if the cache is smaller than the current heatloss possible we know it will not get better
				continue
			}
		}
		cache[currElem] = heatLoss

		if currElem.straight >= minStraight { // if the min straight is over, we need to check if the directions work

			// turn left and right
			left := currElem.dir.Turn(util.Left)
			priorityQueue.Push(QueueEntry{
				pos:      currElem.pos.Move(left, 1),
				dir:      left,
				straight: 1,
			}, heatLoss)

			right := currElem.dir.Turn(util.Right)
			priorityQueue.Push(QueueEntry{
				pos:      currElem.pos.Move(right, 1),
				dir:      right,
				straight: 1,
			}, heatLoss)

		}

		if currElem.straight < maxStraight { // if the maxstraight is not reached go also in that direction
			priorityQueue.Push(QueueEntry{
				pos:      currElem.pos.Move(currElem.dir, 1),
				dir:      currElem.dir,
				straight: currElem.straight + 1,
			}, heatLoss)
		}

	}
	panic("no result found")
}
