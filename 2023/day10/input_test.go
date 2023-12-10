package main

import (
	"testing"

	"github.com/troppes/advent-of-code/util"
)

var testsPart1 = []struct {
	name     string
	input    string
	solution int
}{
	{"e1",
		`.....
.S-7.
.|.|.
.L-J.
.....`, 4},
	{"e1",
		`7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ`, 8},
	{"full", util.ReadInput("input_test.txt"), 80},
}

func TestFullPart1(t *testing.T) {

	for _, test := range testsPart1 {
		t.Run(test.name, func(*testing.T) {
			got := part1(test.input)

			if got != test.solution {
				t.Errorf("Testcase %v failed: got %v and wanted %v", test.name, got, test.solution)
			}
		})
	}
}

var testsPart2 = []struct {
	name     string
	input    string
	solution int
}{
	{"e1",
		`...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`, 4},
	{"e2", `.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...`, 8},
	{"e2", util.ReadInput("input_test.txt"), 10},
}

func TestFullPart2(t *testing.T) {

	for _, test := range testsPart2 {
		t.Run(test.name, func(*testing.T) {
			got := part2(test.input)

			if got != test.solution {
				t.Errorf("Testcase %v failed: got %v and wanted %v", test.name, got, test.solution)
			}
		})
	}
}
