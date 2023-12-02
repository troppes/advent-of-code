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
	{"Game1", `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green`, 1},
	{"Game2", `Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue`, 2},
	{"Game3", `Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red`, 0},
	{"Game4", `Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red`, 0},
	{"Game5", `Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`, 5},
	{"full", util.ReadInput("input_test.txt"), 8},
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
	{"Game1", `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green`, 48},
	{"Game2", `Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue`, 12},
	{"Game3", `Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red`, 1560},
	{"Game4", `Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red`, 630},
	{"Game5", `Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`, 36},
	{"full", util.ReadInput("input_test.txt"), 2286},
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
