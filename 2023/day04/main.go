package main

import (
	"flag"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/troppes/advent-of-code/util"
)

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
	for _, card := range inputArray {

		parts := strings.FieldsFunc(card, func(r rune) bool {
			return r == ':' || r == '|'
		})

		winningPart := util.Filter(strings.Split(parts[1], " "), func(e string) bool {
			return e != ""
		})
		chosenPart := util.Filter(strings.Split(parts[2], " "), func(e string) bool {
			return e != ""
		})

		winning := util.Map(winningPart, func(e string) int {
			no, _ := strconv.Atoi(e)
			return no
		})

		chosen := util.Map(chosenPart, func(e string) int {
			no, _ := strconv.Atoi(e)
			return no
		})

		round := 0
		for _, c := range chosen {
			index := slices.IndexFunc(winning, func(elem int) bool {
				return elem == c
			})
			if index != -1 {
				if round == 0 {
					round = 1
				} else {
					round = round * 2
				}
			}
		}

		solution += round
	}
	return solution
}

func part2(input string) int {
	inputArray := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	cards := make([]int, len(inputArray))

	for i := range cards {
		cards[i] = 1
	}

	for index, card := range inputArray {

		parts := strings.FieldsFunc(card, func(r rune) bool {
			return r == ':' || r == '|'
		})

		winningPart := util.Filter(strings.Split(parts[1], " "), func(e string) bool {
			return e != ""
		})
		chosenPart := util.Filter(strings.Split(parts[2], " "), func(e string) bool {
			return e != ""
		})

		winning := util.Map(winningPart, func(e string) int {
			no, _ := strconv.Atoi(e)
			return no
		})

		chosen := util.Map(chosenPart, func(e string) int {
			no, _ := strconv.Atoi(e)
			return no
		})

		for i := 1; i <= cards[index]; i++ {
			newCards := 0
			for _, c := range chosen {
				index := slices.IndexFunc(winning, func(elem int) bool {
					return elem == c
				})
				if index != -1 {
					newCards++
				}
			}

			for j := 1; j <= newCards; j++ {
				indexToModify := index + j
				if indexToModify >= len(cards) {
					break
				}
				cards[indexToModify]++
			}
		}
	}
	return util.Reduce(cards, func(acc int, val int) int {
		return acc + val
	}, 0)
}
