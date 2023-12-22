package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/troppes/advent-of-code/util"
)

type Pulse struct {
	from  string
	to    string
	pulse int
}

type Module struct {
	moduleType   int
	on           bool
	destinations []string
	inputs       []InputState
}

type InputState struct {
	key   string
	state int
}

const (
	HIGH_PULSE  = 0
	LOW_PULSE   = 1
	FLIP_FLOP   = '%'
	CONJUNCTION = '&'
	BROADCASTER = 'B'
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
	input = strings.ReplaceAll(input, "\r\n", "\n")

	modules := parseModules(input)

	low, high := 0, 0
	for i := 0; i < 1000; i++ { // warm up with 1000 presses
		pulses := pressButton(&modules)
		for _, pulse := range pulses {
			if pulse.pulse == HIGH_PULSE {
				high++
			} else {
				low++
			}
		}
	}
	return low * high
}

func part2(input string) int {
	return 0
}

func parseModules(input string) map[string]Module {

	inputArray := strings.Split(input, "\n")
	modules := make(map[string]Module, len(inputArray))

	for _, line := range inputArray {
		parts := strings.Split(line, " -> ")
		dests := strings.Split(parts[1], ", ")
		moduleType := -1

		key := parts[0]
		if key[0] == '%' || key[0] == '&' {
			moduleType = int(key[0])
			key = key[1:]
		} else {
			moduleType = BROADCASTER
		}

		modules[key] = Module{
			moduleType:   moduleType,
			destinations: dests,
			on:           false,
		}
	}

	// process inputs
	for key, module := range modules {
		for _, dest := range module.destinations {
			if modules[dest].moduleType != CONJUNCTION {
				continue
			}
			c := modules[dest]
			c.inputs = append(modules[dest].inputs, InputState{key, LOW_PULSE})
			modules[dest] = c
		}
	}
	return modules
}

// needed help so used this: https://github.com/JosueMolinaMorales/advent-of-code/blob/main/2023/internal/twenty/day20.go
func pressButton(modules *map[string]Module) []Pulse {
	queue := []Pulse{{"button", "broadcaster", LOW_PULSE}}
	pulses := []Pulse{}

	for len(queue) > 0 { // BFS
		curr := queue[0]
		pulses = append(pulses, curr)
		queue = queue[1:]
		to := (*modules)[curr.to]

		if to.moduleType == FLIP_FLOP {
			if curr.pulse == HIGH_PULSE {
				continue
			}
			to.on = !to.on
			pulse := LOW_PULSE
			if to.on {
				pulse = HIGH_PULSE
			} else {
				pulse = LOW_PULSE
			}
			(*modules)[curr.to] = to
			for _, out := range to.destinations { // broadcast
				queue = append(queue, Pulse{curr.to, out, pulse})
			}
		} else if to.moduleType == CONJUNCTION {
			m := (*modules)[curr.to]
			m.inputs = util.Map(m.inputs, func(input InputState) InputState {
				if input.key == curr.from {
					input.state = curr.pulse
				}
				return input
			})
			(*modules)[curr.to] = m
			allHigh := util.Every(m.inputs, func(input InputState) bool {
				return input.state == HIGH_PULSE
			})
			pulse := HIGH_PULSE
			if allHigh {
				pulse = LOW_PULSE
			}
			for _, out := range to.destinations {
				queue = append(queue, Pulse{curr.to, out, pulse})
			}
		} else {
			for _, out := range to.destinations {
				queue = append(queue, Pulse{curr.to, out, LOW_PULSE})
			}
		}
	}
	return pulses
}
