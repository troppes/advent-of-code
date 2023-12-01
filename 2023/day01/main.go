package main

import (
	"flag"
	"fmt"
	"regexp"
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

	numbersRegEXP := regexp.MustCompile(`\d`)
	solution := 0
	for _, s := range inputArray {
		numbers := numbersRegEXP.FindAllString(s, -1)
		num, _ := strconv.Atoi(numbers[0] + numbers[len(numbers)-1])
		solution += num
	}
	return solution
}

func part2(input string) int {
	// The weirdest hack of all time, the numbers can overlap
	replacer := strings.NewReplacer(
		"one", "o1e",
		"two", "t2o",
		"three", "t3e",
		"four", "f4r",
		"five", "f5e",
		"six", "s6x",
		"seven", "s7n",
		"eight", "e8t",
		"nine", "n9e")
	newInput := replacer.Replace(input)
	newInput = replacer.Replace(newInput)
	return part1(newInput)
}
