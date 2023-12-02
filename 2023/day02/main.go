package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/troppes/advent-of-code/util"
)

var limits = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
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
		ans := part2(util.ReadInput("./input.txt"))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	inputArray := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
	solution := 0

	for _, s := range inputArray {
		games := strings.FieldsFunc(s, func(r rune) bool {
			return r == ':' || r == ';'
		})
		currentGame, games := strings.Split(games[0], " "), games[1:]
		gameId := currentGame[1]

		gamePossible := true
		for _, round := range games {
			if !isGamePossible(round) {
				gamePossible = false
			}
		}
		if gamePossible {
			num, _ := strconv.Atoi(gameId)
			solution += num
		}
	}
	return solution
}

func part2(input string) int {
	inputArray := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
	solution := 0

	for _, s := range inputArray {
		games := strings.FieldsFunc(s, func(r rune) bool {
			return r == ':' || r == ';'
		})

		var currentNumber = map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		games = games[1:]

		for _, round := range games {
			colors := strings.Split(round, ",")
			for _, colorCombination := range colors {
				c := strings.Split(colorCombination, " ")
				num, _ := strconv.Atoi(c[1])
				if num > currentNumber[c[2]] {
					currentNumber[c[2]] = num
				}
			}
		}
		solution += currentNumber["red"] * currentNumber["green"] * currentNumber["blue"]
	}
	return solution
}

func isGamePossible(round string) bool {
	possible := true
	colors := strings.Split(round, ",")

	for _, c := range colors {
		colorCombination := strings.Split(c, " ")
		num, _ := strconv.Atoi(colorCombination[1])
		if num > limits[colorCombination[2]] {
			possible = false
		}
	}
	return possible

}
