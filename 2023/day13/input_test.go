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
	{"line1", `???.### 1,1,3`, 1},
	{"line2", `.??..??...?##. 1,1,3`, 4},
	{"line3", `?#?#?#?#?#?#?#? 1,3,1,6`, 1},
	{"line4", `????.#...#... 4,1,1`, 1},
	{"line5", `????.######..#####. 1,6,5`, 4},
	{"line6", `?###???????? 3,2,1`, 10},
	{"full", util.ReadInput("input_test.txt"), 21},
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
	{"line1", `???.### 1,1,3`, 1},
	{"line2", `.??..??...?##. 1,1,3`, 16384},
	{"line3", `?#?#?#?#?#?#?#? 1,3,1,6`, 1},
	{"line4", `????.#...#... 4,1,1`, 16},
	{"line5", `????.######..#####. 1,6,5`, 2500},
	{"line6", `?###???????? 3,2,1`, 506250},
	{"full", util.ReadInput("input_test.txt"), 525152},
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
