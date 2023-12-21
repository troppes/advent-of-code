package main

import (
	"flag"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/troppes/advent-of-code/util"
)

type Rule struct {
	partType  string
	operation string
	value     int
	next      string
}

type Part struct {
	x        int
	m        int
	a        int
	s        int
	accepted bool
}

type Workflows = map[string]util.List[Rule]

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()

	input := util.ReadInput("input_test.txt")

	if part == 1 {
		ans := part1(input)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	workflows, parts := parseInput(input)
	return evaluateParts(parts, workflows)
}

func part2(input string) int {
	return 0
}

func parseInput(input string) (Workflows, []Part) {
	inputArray := strings.Split(input, "\n\n")

	workflowList := strings.Split(inputArray[0], "\n")
	partsList := strings.Split(inputArray[1], "\n")

	parts := make([]Part, len(partsList))
	workflows := make(map[string]util.List[Rule], len(workflowList))

	for i, part := range partsList {
		parts[i] = parsePart(part)
	}

	for _, workflow := range workflowList {
		name, workflow := parseWorkflow(workflow)
		workflows[name] = workflow
	}

	return workflows, parts

}

func parseWorkflow(workflowString string) (string, util.List[Rule]) {
	re := regexp.MustCompile(`^([^{}]+){([^{}]+)}$`)

	list := util.List[Rule]{}
	matches := re.FindStringSubmatch(workflowString)

	if len(matches) != 3 {
		panic("Invalid input string")
	}

	beginning := matches[1]
	content := matches[2]

	ruleRegex := regexp.MustCompile(`^([a-zA-Z])([<>])(\d+):([a-zA-Z]+)$`)
	for _, rule := range strings.Split(content, ",") {

		// Find submatches
		matches := ruleRegex.FindStringSubmatch(rule)

		if len(matches) == 5 {
			num, err := strconv.Atoi(matches[3])
			if err != nil {
				panic("Can not convert to Number")
			}
			list.Insert(Rule{partType: matches[1], operation: matches[2], value: num, next: matches[4]})
		} else if len(matches) == 0 { // catch if last elem just is something to go to
			list.Insert(Rule{operation: "=", next: rule})
		} else {
			panic("Not sufficient parts for a rule")
		}

	}
	return beginning, list
}

func parsePart(partString string) Part {
	re := regexp.MustCompile(`(\d+)`)
	matches := re.FindAllString(partString, -1)

	if len(matches) != 4 {
		panic("No sufficient parts")
	}

	// Convert strings to integers
	numbers := make([]int, len(matches))
	for i, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			panic("Can not convert to Number")
		}
		numbers[i] = num
	}

	return Part{
		x:        numbers[0],
		m:        numbers[1],
		a:        numbers[2],
		s:        numbers[3],
		accepted: false,
	}
}

func evaluateParts(parts []Part, workflows map[string]util.List[Rule]) int {

	count := 0
	for _, part := range parts {
		currentId := "in"
		for currentId != "A" && currentId != "R" {
			currentWorkflow := workflows[currentId].Head
			for currentWorkflow.Next != nil {
				if evaluateRule(part, currentWorkflow.Info) {
					break
				}
				currentWorkflow = currentWorkflow.Next
			}
			currentId = currentWorkflow.Info.next
		}
		if currentId == "A" {
			count += part.x + part.m + part.a + part.s
		}
	}

	return count
}

func evaluateRule(part Part, rule Rule) bool {

	if rule.operation == "=" { // equal rules are always true
		return true
	}

	val := reflect.ValueOf(part) // crazy evil, but it works
	field := val.FieldByName(rule.partType)

	if field.IsValid() {
		curr := field.Int()
		if rule.operation == "<" {
			return curr < int64(rule.value)
		} else {
			return curr > int64(rule.value)
		}
	} else {
		panic("Field not found")
	}
}
