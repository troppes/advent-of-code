package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"

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

	lookupMap := parseGrid(input)

	for y := 0; y < len(inputArray); y++ {
		line := inputArray[y]

		re := regexp.MustCompile("[0-9]+")

		matches := re.FindAllStringSubmatchIndex(line, -1)

		for _, match := range matches {
			startIndex := match[0]
			endIndex := match[1]
			numberString := line[startIndex:endIndex]

			isPart := false
			for x := startIndex; x < endIndex; x++ {
				if checkSurroundingsForSpecial(util.Coordinate{X: x, Y: y}, lookupMap) {
					isPart = true
				}
			}

			if isPart {
				num, _ := strconv.Atoi(numberString)
				solution += num
			}

		}
	}

	return solution
}

func part2(input string) int {
	inputArray := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
	solution := 0

	lookupMap := parseGrid(input)

	for y := 0; y < len(inputArray); y++ {
		line := inputArray[y]

		for x, char := range line {
			if char != '.' && !unicode.IsDigit(char) {
				numbers := checkSurroundingsForNumbers(util.Coordinate{X: x, Y: y}, lookupMap)

				if len(numbers) == 2 {
					solution += (numbers[0] * numbers[1])
				}
			}
		}
	}

	return solution
}

func parseGrid(input string) map[util.Coordinate]rune {
	inputArray := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	coordinates := make(map[util.Coordinate]rune)

	for y, line := range inputArray {
		for x, character := range line {
			coordinates[util.Coordinate{X: x, Y: y}] = character
		}
	}

	return coordinates
}

func checkSurroundingsForSpecial(coordinate util.Coordinate, lookUp map[util.Coordinate]rune) bool {

	for _, surrCoord := range []util.Coordinate{
		{X: -1, Y: 0}, {X: 1, Y: 0}, // X
		{X: 0, Y: -1}, {X: 0, Y: 1}, // Y
		{X: -1, Y: 1}, {X: 1, Y: 1}, // diag Top
		{X: -1, Y: -1}, {X: 1, Y: -1}, // diag bottom
	} {
		coordToCheck := util.Coordinate{X: coordinate.X + surrCoord.X, Y: coordinate.Y + surrCoord.Y}
		if lookUp[coordToCheck] != 0 && lookUp[coordToCheck] != '.' && !unicode.IsDigit(lookUp[coordToCheck]) {
			return true
		}
	}

	return false
}

func checkSurroundingsForNumbers(coordinate util.Coordinate, lookUp map[util.Coordinate]rune) []int {

	var surroundingNumbers = make(map[util.Coordinate]int)

	for _, surrCoord := range []util.Coordinate{
		{X: -1, Y: 0}, {X: 1, Y: 0}, // X
		{X: 0, Y: -1}, {X: 0, Y: 1}, // Y
		{X: -1, Y: 1}, {X: 1, Y: 1}, // diag Top
		{X: -1, Y: -1}, {X: 1, Y: -1}, // diag bottom
	} {
		coordToCheck := util.Coordinate{X: coordinate.X + surrCoord.X, Y: coordinate.Y + surrCoord.Y}
		if lookUp[coordToCheck] != 0 && unicode.IsDigit(lookUp[coordToCheck]) {
			// check both directions if found

			numString := string(lookUp[coordToCheck])
			xPos := coordToCheck.X + 1
			xNeg := coordToCheck.X - 1
			for unicode.IsDigit(lookUp[util.Coordinate{Y: coordToCheck.Y, X: xPos}]) {
				numString += string(lookUp[util.Coordinate{Y: coordToCheck.Y, X: xPos}])
				xPos++
			}
			for unicode.IsDigit(lookUp[util.Coordinate{Y: coordToCheck.Y, X: xNeg}]) {
				numString = string(lookUp[util.Coordinate{Y: coordToCheck.Y, X: xNeg}]) + numString
				xNeg--
			}

			num, _ := strconv.Atoi(numString)

			surroundingNumbers[util.Coordinate{Y: coordToCheck.Y, X: xNeg}] = num // take first index to catch repeating number

		}
	}

	var numbers []int
	for _, value := range surroundingNumbers {
		numbers = append(numbers, value)
	}

	return numbers
}
