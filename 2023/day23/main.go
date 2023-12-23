package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/oleiade/lane/v2" // queue
	"github.com/troppes/advent-of-code/util"
)

type QueueEntry struct {
	coord   util.Coordinate
	steps   int
	journey map[util.Coordinate]bool
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
	grid, target := parseGrid(input, false)
	return findLongestJourney(grid, target)
}

func part2(input string) int {
	grid, target := parseGrid(input, true)
	return findLongestJourney(grid, target)
}

func parseGrid(input string, onlyDot bool) (map[util.Coordinate]rune, util.Coordinate) {
	inputArray := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	coordinates := make(map[util.Coordinate]rune)

	for y, line := range inputArray {
		for x, character := range line {
			if onlyDot && (character == '<' || character == '>' || character == '^' || character == 'v') {
				character = '.'
			}
			coordinates[util.Coordinate{X: x, Y: y}] = character

		}
	}

	return coordinates, util.Coordinate{Y: len(inputArray) - 1, X: len(inputArray[0]) - 2}
}

func findLongestJourney(grid map[util.Coordinate]rune, target util.Coordinate) int {

	queue := lane.NewQueue[QueueEntry]()

	queue.Enqueue(QueueEntry{
		coord:   util.Coordinate{Y: 0, X: 1},
		steps:   0,
		journey: map[util.Coordinate]bool{{Y: 0, X: 1}: true},
	})

	journeys := make(map[util.Coordinate]int)

	for queue.Size() > 0 {
		currElem, _ := queue.Dequeue()

		if steps, ok := journeys[currElem.coord]; !ok || currElem.steps > steps {
			journeys[currElem.coord] = currElem.steps
			for _, neighbour := range neighbours(grid, currElem) {
				if !currElem.journey[neighbour.coord] { // only enqueue if the coordinate is not in the path
					neighbour.journey = make(map[util.Coordinate]bool)
					for k, v := range currElem.journey {
						neighbour.journey[k] = v
					}
					neighbour.journey[neighbour.coord] = true
					queue.Enqueue(neighbour)
				}
			}
		}
	}
	return journeys[target]
}

func neighbours(grid map[util.Coordinate]rune, currElem QueueEntry) []QueueEntry {
	newPos := currElem.coord
	switch grid[currElem.coord] {
	case 'v':
		newPos.Y++
		return []QueueEntry{{coord: newPos, steps: currElem.steps + 1}}
	case '^':
		newPos.Y--
		return []QueueEntry{{coord: newPos, steps: currElem.steps + 1}}
	case '>':
		newPos.X++
		return []QueueEntry{{coord: newPos, steps: currElem.steps + 1}}
	case '<':
		newPos.X--
		return []QueueEntry{{coord: newPos, steps: currElem.steps + 1}}
	}

	var neighbours []QueueEntry

	for _, direction := range []util.Direction{util.Up, util.Down, util.Left, util.Right} {
		newPos = currElem.coord.Move(direction, 1)
		if character, exists := grid[newPos]; !exists || character == '#' { // oob check
			continue
		}
		neighbours = append(neighbours, QueueEntry{coord: newPos, steps: currElem.steps + 1})

	}
	return neighbours
}
