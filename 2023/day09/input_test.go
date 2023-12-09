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
	{"l1", "0 3 6 9 12 15", 18},
	{"l1", "1 3 6 10 15 21", 28},
	{"l1", "10 13 16 21 30 45", 68},
	{"full", util.ReadInput("input_test.txt"), 114},
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
	{"l1", "0 3 6 9 12 15", -3},
	{"l1", "1 3 6 10 15 21", 0},
	{"l1", "10 13 16 21 30 45", 5},
	{"full", util.ReadInput("input_test.txt"), 2},
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
