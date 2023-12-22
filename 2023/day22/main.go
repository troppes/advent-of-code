package main

import (
	"flag"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/troppes/advent-of-code/util"
)

type Brick struct {
	x1 int
	y1 int
	z1 int
	x2 int
	y2 int
	z2 int
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
	input = strings.ReplaceAll(input, "\r\n", "\n")
	bricks := parseBricks(input)

	slices.SortFunc(bricks, func(a, b Brick) int {
		return min(a.z1, a.z2) - min(b.z1, b.z2)
	})

	moveBricks(bricks)

	count := 0

	for i := range bricks {
		tmp := make([]Brick, len(bricks))
		copy(tmp, bricks)
		tmp = util.RemoveAtIndex(tmp, i)
		changes := moveBricks(tmp)

		if changes == 0 {
			count++
		}
	}

	return count
}

func part2(input string) int {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	bricks := parseBricks(input)

	slices.SortFunc(bricks, func(a, b Brick) int {
		return min(a.z1, a.z2) - min(b.z1, b.z2)
	})

	moveBricks(bricks)

	count := 0

	for i := range bricks {
		tmp := make([]Brick, len(bricks))
		copy(tmp, bricks)
		tmp = util.RemoveAtIndex(tmp, i)
		count += moveBricks(tmp)
	}

	return count
}

func parseBricks(input string) []Brick {

	inputArray := strings.Split(input, "\n")
	bricks := make([]Brick, len(inputArray))

	re := regexp.MustCompile(`\d+`)
	for i, brick := range inputArray {

		// Find all number matches in the input string
		matches := re.FindAllString(brick, -1)

		intMatches := util.Map(matches, func(s string) int {
			num, err := strconv.Atoi(s)
			if err != nil {
				panic("Not a valid number")
			}
			return num
		})

		bricks[i] = Brick{x1: intMatches[0], y1: intMatches[1], z1: intMatches[2], x2: intMatches[3], y2: intMatches[4], z2: intMatches[5]}
	}
	return bricks
}

// was not able to figure this one out by myself, used https://github.com/dannyvankooten/advent-of-code/blob/main/2023/22-sand-slabs/main.go to help me
func moveBricks(bricks []Brick) int {
	blocksMoved := 0

nextbrick:
	for i := range bricks {
		a := &bricks[i] // use pointer so we can change the values
		moved := false

		for a.z1 > 1 {
			for j := i - 1; j >= 0; j-- { // look at all blocks bleow this one bc sort
				b := bricks[j]

				if (a.z2-1) >= b.z1 && (a.z1-1) <= b.z2 && // check if we are intersecting when moving down
					a.x2 >= b.x1 && a.x1 <= b.x2 &&
					a.y2 >= b.y1 && a.y1 <= b.y2 {
					continue nextbrick // if so stop and move to next brick
				}
			}

			if !moved { // if we moved we add it to the needed to change list
				blocksMoved++
				moved = true
			}

			// move the block down
			a.z1 -= 1
			a.z2 -= 1
		}

	}
	return blocksMoved
}
