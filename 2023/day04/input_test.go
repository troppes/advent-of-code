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
	{"Game1", `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53`, 8},
	{"Game2", `Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19`, 2},
	{"Game3", `Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1`, 2},
	{"Game4", `Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83`, 1},
	{"Game5", `Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36`, 0},
	{"Game5", `Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`, 0},
	{"full", util.ReadInput("input_test.txt"), 13},
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
	{"full", util.ReadInput("input_test.txt"), 30},
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
