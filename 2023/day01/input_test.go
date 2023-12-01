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
	{"line1", `1abc2`, 12},
	{"line2", `pqr3stu8vwx`, 38},
	{"line3", `a1b2c3d4e5f`, 15},
	{"line4", `treb7uchet`, 77},
	{"line5", `12abc2`, 12},
	{"full", util.ReadInput("input_test.txt"), 142},
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
	{"line1", `two1nine`, 29},
	{"line2", `eightwothree`, 83},
	{"line3", `abcone2threexyz`, 13},
	{"line4", `xtwone3four`, 24},
	{"line5", `4nineeightseven2`, 42},
	{"line5", `zoneight234`, 14},
	{"line5", `7pqrstsixteen`, 76},
	{"line6", `two1nineeightnine`, 29},
	{"reallyForDayOne?", `eighthree`, 83},
	{"reallyForDayOne?P2", `sevenine`, 79},
	{"full", util.ReadInput("input_test2.txt"), 281},
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
