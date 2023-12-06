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
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	inputArray := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	// Define the regular expression to find integers
	re := regexp.MustCompile(`\d+`)

	times := util.Map(re.FindAllString(inputArray[0], -1), func(s string) int {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return num
	})

	distances := util.Map(re.FindAllString(inputArray[1], -1), func(s string) int {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return num
	})

	sol := 1
	for i, time := range times {
		var results []int
		for held := 0; held < time; held++ {
			results = append(results, calcSpeed(held, time))
		}

		curr := 0
		for _, result := range results {
			if result > distances[i] {
				curr++
			}
		}
		sol *= curr

	}

	return sol
}

func part2(input string) int {
	return part1(strings.ReplaceAll(input, " ", ""))
}

func calcSpeed(holdtime int, fullTime int) int {
	return holdtime * (fullTime - holdtime)
}
