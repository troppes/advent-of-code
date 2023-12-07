package main

import (
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/troppes/advent-of-code/util"
)

type Round struct {
	value int
	hand  string
	rank  int
}

var cardValues = map[string]int{
	"A": 12,
	"K": 11,
	"Q": 10,
	"J": 9,
	"T": 8,
	"9": 7,
	"8": 6,
	"7": 5,
	"6": 4,
	"5": 3,
	"4": 2,
	"3": 1,
	"2": 0,
}

var cardValuesPart2 = map[string]int{
	"A": 12,
	"K": 11,
	"Q": 10,
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
	"J": 0,
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
	inputArray := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	var rounds []Round

	for _, round := range inputArray {
		round := strings.Split(round, " ")

		num, err := strconv.Atoi(round[1])
		if err != nil {
			panic(err)
		}

		rank := evalHand(round[0])
		rounds = append(rounds, Round{value: num, hand: round[0], rank: rank})
	}

	sort.Slice(rounds, func(i, j int) bool {

		x := rounds[i]
		y := rounds[j]

		if x.rank == y.rank {
			cardX := strings.Split(x.hand, "")
			cardY := strings.Split(y.hand, "")
			for i := range cardX {
				if cardValues[cardX[i]] < cardValues[cardY[i]] {
					return true
				} else if cardValues[cardX[i]] > cardValues[cardY[i]] {
					return false
				}
			}
		}
		return x.rank < y.rank
	})

	sol := 0
	for i, round := range rounds {
		sol += (i + 1) * round.value
	}

	return sol
}

func part2(input string) int {
	inputArray := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	var rounds []Round

	for _, round := range inputArray {
		round := strings.Split(round, " ")

		num, err := strconv.Atoi(round[1])
		if err != nil {
			panic(err)
		}

		rank := evalHandPart2(round[0])
		rounds = append(rounds, Round{value: num, hand: round[0], rank: rank})
	}

	sort.Slice(rounds, func(i, j int) bool {

		x := rounds[i]
		y := rounds[j]

		if x.rank == y.rank {
			cardX := strings.Split(x.hand, "")
			cardY := strings.Split(y.hand, "")
			for i := range cardX {
				if cardValuesPart2[cardX[i]] < cardValuesPart2[cardY[i]] {
					return true
				} else if cardValuesPart2[cardX[i]] > cardValuesPart2[cardY[i]] {
					return false
				}
			}
		}
		return x.rank < y.rank
	})

	sol := 0
	for i, round := range rounds {
		sol += (i + 1) * round.value
	}

	return sol
}

func evalHand(hand string) int {

	rank := -1

	labels := make(map[rune]int)

	for _, char := range hand {
		labels[char] += 1
	}

	var counts []int
	for _, v := range labels {
		counts = append(counts, v)
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})

	if counts[0] == 5 {
		rank = 6
	} else if counts[0] == 4 {
		rank = 5
	} else if counts[0] == 3 && counts[1] == 2 {
		rank = 4
	} else if counts[0] == 3 {
		rank = 3
	} else if counts[0] == 2 && counts[1] == 2 {
		rank = 2
	} else if counts[0] == 2 {
		rank = 1
	} else {
		rank = 0
	}

	return rank
}

func evalHandPart2(hand string) int {

	rank := -1

	labels := make(map[string]int)
	jokers := 0

	for _, char := range strings.Split(hand, "") {
		if char == "J" {
			jokers++
		} else {
			labels[char] += 1
		}
	}

	var counts []int
	for _, v := range labels {
		counts = append(counts, v)
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})

	if len(counts) > 0 {
		counts[0] += jokers
	} else {
		counts = append(counts, jokers)
	}

	if counts[0] == 5 {
		rank = 6
	} else if counts[0] == 4 {
		rank = 5
	} else if counts[0] == 3 && counts[1] == 2 {
		rank = 4
	} else if counts[0] == 3 {
		rank = 3
	} else if counts[0] == 2 && counts[1] == 2 {
		rank = 2
	} else if counts[0] == 2 {
		rank = 1
	} else {
		rank = 0
	}

	return rank
}
