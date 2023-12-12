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

	count := 0
	for _, line := range inputArray {
		data := strings.Split(line, " ")
		pattern := data[0]
		groups := util.Map(strings.Split(data[1], ","), func(s string) int {
			num, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			return num
		})

		var cache [][]int
		for i := 0; i < len(pattern); i++ {
			cache = append(cache, make([]int, len(groups)+1))
			for j := 0; j < len(groups)+1; j++ {
				cache[i][j] = -1
			}
		}

		count += findCombinations(0, 0, pattern, groups, cache)
	}

	return count
}

func part2(input string) int {
	inputArray := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	count := 0

	for _, line := range inputArray {
		data := strings.Split(line, " ")
		pattern := data[0]
		groups := util.Map(strings.Split(data[1], ","), func(s string) int {
			num, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			return num
		})

		unfoldedPattern := unfoldPattern(pattern)
		unfoldedGroups := unfoldGroups(groups)

		var cache [][]int
		for i := 0; i < len(unfoldedPattern); i++ {
			cache = append(cache, make([]int, len(unfoldedGroups)+1))
			for j := 0; j < len(unfoldedGroups)+1; j++ {
				cache[i][j] = -1
			}
		}

		count += findCombinations(0, 0, unfoldedPattern, unfoldedGroups, cache)
	}

	return count
}

// needed help for this function, based on: https://github.com/ayoubzulfiqar/advent-of-code/blob/main/Go/Day12/part_2.go
func findCombinations(indexP, indexG int, pattern string, groups []int, cache [][]int) int {
	if indexP >= len(pattern) {
		if indexG >= len(groups) {
			return 1
		}
		return 0
	}

	if cache[indexP][indexG] != -1 {
		return cache[indexP][indexG]
	}

	permutations := 0
	currByte := pattern[indexP]

	if currByte == '.' {
		permutations = findCombinations(indexP+1, indexG, pattern, groups, cache)
	} else {
		if currByte == '?' { // treats the ? as .
			permutations += findCombinations(indexP+1, indexG, pattern, groups, cache)
		}
		// otherwise treat questionmark as # or handle #
		if indexG < len(groups) {
			count := 0
			countNeeded := groups[indexG]
			for i := indexP; i < len(pattern); i++ {
				// if first the count is higher than the count needed, then it is invalid
				// or the pattern is not a #/?, which ends the group as well
				// or we already have the count and the next one is a ? that assumes the role of .
				if count > countNeeded || pattern[i] == '.' || count == countNeeded && pattern[i] == '?' {
					break
				}
				count += 1
			}

			if count == countNeeded {
				// check if we have reached the end of the word and the next one is not a #,
				// since we need a . or a ? that assumes the role of . to continue
				if indexP+count < len(pattern) && pattern[indexP+count] != '#' {
					permutations += findCombinations(indexP+count+1, indexG+1, pattern, groups, cache)
				} else {
					permutations += findCombinations(indexP+count, indexG+1, pattern, groups, cache)
				}
			}
		}
	}

	cache[indexP][indexG] = permutations
	return permutations
}

func unfoldGroups(groups []int) []int {
	newGroups := make([]int, len(groups)*5)
	for i := 0; i < 5; i++ {
		for j := 0; j < len(groups); j++ {
			newGroups[i*len(groups)+j] = groups[j]
		}
	}
	return newGroups
}

func unfoldPattern(pattern string) string {
	var sb strings.Builder
	for i := 0; i < 5; i++ {
		for j := 0; j < len(pattern); j++ {
			sb.WriteByte(pattern[j])
		}
		if i < 4 {
			sb.WriteRune('?')

		}
	}
	return sb.String()
}
