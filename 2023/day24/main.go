package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	// queue
	"github.com/troppes/advent-of-code/util"
)

type Coord3D struct {
	x int
	y int
	z int
}

type HailStone struct {
	pos Coord3D
	vel Coord3D
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

	hailstones := parse(input)
	counter := 0

	start := 200000000000000
	end := 400000000000000

	for i := 0; i < len(hailstones); i++ {
		for j := i + 1; j < len(hailstones); j++ {
			h1 := hailstones[i]
			h2 := hailstones[j]

			if intersect(h1, h2, start, end) {
				counter++
			}
		}
	}
	return counter
}

func part2(input string) int {
	return part1(input)
}

func parse(input string) []HailStone {
	inputArray := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	hailstones := make([]HailStone, len(inputArray))

	for i, line := range inputArray {
		// 19, 13, 30 @ -2,  1, -2

		parts := strings.Split(line, " @ ")
		pos := parseCoord3D(parts[0])
		vel := parseCoord3D(parts[1])
		hailstones[i] = HailStone{pos: pos, vel: vel}
	}

	return hailstones
}

func parseCoord3D(coordString string) Coord3D {
	parts := strings.Split(coordString, ", ")
	coord := util.Map(parts, func(s string) int {
		num, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			panic("Not a valid number")
		}
		return num
	})

	return Coord3D{x: coord[0], y: coord[1], z: coord[2]}
}

func intersect(hail1 HailStone, hail2 HailStone, start int, end int) bool {

	// parallel
	if hail1.vel.x*hail2.vel.y-hail1.vel.y*hail2.vel.x == 0 {
		// check if lines are the same
		return hail1.vel.x*(hail1.pos.y-hail2.pos.y)+hail1.vel.y*(hail2.pos.x-hail1.pos.x) == 0
	}

	// find intersect points
	tangent2 := (hail1.vel.x*(hail1.pos.y-hail2.pos.y) + hail1.vel.y*(hail2.pos.x-hail1.pos.x)) / (hail1.vel.x*hail2.vel.y - hail1.vel.y*hail2.vel.x)
	tangent1 := (hail2.pos.x - hail1.pos.x + hail2.vel.x*tangent2) / hail1.vel.x

	// if we hit before it starts
	if tangent1 < 0 || tangent2 < 0 {
		return false
	}

	// calc coordinates
	intersectX := hail2.pos.x + hail2.vel.x*tangent2
	intersectY := hail2.pos.y + hail2.vel.y*tangent2

	// check with are within the parameters
	return (start <= intersectX && intersectX <= end) &&
		(start <= intersectY && intersectY <= end)
}
