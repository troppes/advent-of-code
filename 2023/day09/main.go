package main

import (
	"flag"
	"fmt"
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
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	inputArray := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	var solution []int

	for _, line := range inputArray {

		numbers := util.Map(strings.Split(line, " "), func(s string) int {
			num, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			return num
		})

		history := calculateHistory(numbers)
		extrapolatedData := extrapolateHistory(history)

		solution = append(solution, extrapolatedData)

	}

	return util.Reduce(solution, func(acc int, val int) int {
		return acc + val
	}, 0)
}

func part2(input string) int {
	inputArray := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	var solution []int

	for _, line := range inputArray {

		numbers := util.Map(strings.Split(line, " "), func(s string) int {
			num, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			return num
		})

		history := calculateHistory(numbers)
		extrapolatedData := extrapolateHistoryBackwards(history)

		solution = append(solution, extrapolatedData)

	}

	return util.Reduce(solution, func(acc int, val int) int {
		return acc + val
	}, 0)
}

func calculateHistory(data []int) [][]int {

	var history [][]int

	history = append(history, data)

	// a bit weird, but basically if the array consits of only zeroes we stop
	for util.Reduce(history[len(history)-1], func(acc int, val int) int {
		return acc + val
	}, 0) != 0 {
		var historyLine []int
		currentHistory := history[len(history)-1]

		// -1 to have the second to last element as last
		for i := 0; i < len(currentHistory)-1; i++ {
			historyLine = append(historyLine, currentHistory[i+1]-currentHistory[i])
		}

		history = append(history, historyLine)
	}

	return history
}

func extrapolateHistory(history [][]int) int {

	// append the zero
	history[len(history)-1] = append(history[len(history)-1], 0)

	for i := len(history) - 2; i >= 0; i-- {
		currentLast := history[i][len(history[i])-1]
		belowLast := history[i+1][len(history[i+1])-1]
		history[i] = append(history[i], currentLast+belowLast)
	}

	return history[0][len(history[0])-1]
}

func extrapolateHistoryBackwards(history [][]int) int {

	// prepend the zero
	history[len(history)-1] = append([]int{0}, history[len(history)-1]...)

	for i := len(history) - 2; i >= 0; i-- {
		currentFirst := history[i][0]
		belowFirst := history[i+1][0]
		history[i] = append([]int{currentFirst - belowFirst}, history[i]...)
	}

	return history[0][0]
}
