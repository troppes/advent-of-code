package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/troppes/advent-of-code/util"
)

type Lense struct {
	focalLength int
	label       string
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
	// just to be sure remove all newlines
	input = strings.ReplaceAll(strings.ReplaceAll(input, "\r\n", "\n"), "\n", "")
	inputArray := strings.Split(input, ",")

	sum := 0
	for _, data := range inputArray {
		sum += createHash([]byte(data))
	}
	return sum
}

func part2(input string) int {
	// just to be sure remove all newlines
	input = strings.ReplaceAll(strings.ReplaceAll(input, "\r\n", "\n"), "\n", "")
	inputArray := strings.Split(input, ",")
	sum := 0

	boxes := make([]util.List[Lense], 256)
	for i := range boxes {
		boxes[i] = util.List[Lense]{} // Initialize each list

		// Set the matcher for each list in the boxes slice
		boxes[i].SetMatcher(func(t1, t2 *Lense) bool {
			return t1.label == t2.label
		})
	}

	for _, data := range inputArray {
		re := regexp.MustCompile(`(\b\w+)([-=])(\d*)`)
		match := re.FindStringSubmatch(data) // 0 = full string / 1 = toHash / 2 = sign / 3 = number
		// checked with regex, there are not negative ones with numbers attached
		hash := createHash([]byte(match[1]))

		if match[2] == "=" {
			num, err := strconv.Atoi(match[3])
			if err != nil {
				panic(err)
			}
			boxes[hash].InsertAndReplace(Lense{focalLength: num, label: match[1]})
		} else {
			boxes[hash].Delete(Lense{focalLength: -1, label: match[1]})
		}
	}

	for j, box := range boxes {
		if !box.Empty() {
			fmt.Print("Box ", j, ": ")
			box.Print()
		}
	}

	for i, box := range boxes {
		indexList := 1
		curr := box.Head
		for curr != nil {
			sum += (i + 1) * indexList * curr.Info.focalLength
			curr = curr.Next
			indexList++
		}
	}
	return sum
}

func createHash(bytes []byte) int {

	curr := 0

	for _, b := range bytes {
		curr += int(b)
		curr *= 17
		curr %= 256
	}

	return curr
}
