package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/troppes/advent-of-code/util"
)

type GardenMap struct {
	from int
	to   int
	rng  int
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
	const MaxUint = ^uint(0)
	currMin := int(MaxUint >> 1) // maxint
	sections := strings.Split(input, "\n\n")

	sections = util.Map(sections, func(s string) string { // trim everthing before the numbers
		return strings.TrimLeft(s, "-: \nabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	})
	seeds := util.Map(strings.Split(sections[0], " "), func(s string) int {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return num
	})

	seedToSoil := createMap(sections[1])
	soilToFert := createMap(sections[2])
	fertToWater := createMap(sections[3])
	waterToLight := createMap(sections[4])
	lightToTemp := createMap(sections[5])
	tempToHumid := createMap(sections[6])
	humidToLoc := createMap(sections[7])

	for _, seed := range seeds {
		loc := findNewLocation(seed, seedToSoil)
		loc = findNewLocation(loc, soilToFert)
		loc = findNewLocation(loc, fertToWater)
		loc = findNewLocation(loc, waterToLight)
		loc = findNewLocation(loc, lightToTemp)
		loc = findNewLocation(loc, tempToHumid)
		loc = findNewLocation(loc, humidToLoc)
		if currMin > loc {
			currMin = loc
		}
	}

	return currMin
}

func part2(input string) int {

	const MaxUint = ^uint(0)
	currMin := int(MaxUint >> 1) // maxint
	sections := strings.Split(input, "\n\n")

	sections = util.Map(sections, func(s string) string { // trim everthing before the numbers
		return strings.TrimLeft(s, "-: \nabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	})
	seeds := util.Map(strings.Split(sections[0], " "), func(s string) int {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return num
	})

	seedToSoil := createMap(sections[1])
	soilToFert := createMap(sections[2])
	fertToWater := createMap(sections[3])
	waterToLight := createMap(sections[4])
	lightToTemp := createMap(sections[5])
	tempToHumid := createMap(sections[6])
	humidToLoc := createMap(sections[7])

	for i := 0; i < len(seeds); i += 2 {
		for seed := seeds[i]; seed < (seeds[i] + seeds[i+1]); seed++ {

			loc := findNewLocation(seed, seedToSoil)
			loc = findNewLocation(loc, soilToFert)
			loc = findNewLocation(loc, fertToWater)
			loc = findNewLocation(loc, waterToLight)
			loc = findNewLocation(loc, lightToTemp)
			loc = findNewLocation(loc, tempToHumid)
			loc = findNewLocation(loc, humidToLoc)

			if currMin > loc {
				currMin = loc
			}

		}
	}
	return currMin
}

func createMap(input string) []GardenMap {
	lines := strings.Split(input, "\n")

	var gardenMaps []GardenMap
	for _, line := range lines {
		gardenMaps = append(gardenMaps, processLine(line))
	}

	return gardenMaps
}

func processLine(line string) GardenMap {
	gMap := new(GardenMap)

	split := util.Map(strings.Split(line, " "), func(x string) int {
		num, err := strconv.Atoi(x)
		if err != nil {
			panic(err)
		}
		return num
	})

	gMap.from = split[1]
	gMap.to = split[0]
	gMap.rng = split[2]

	return *gMap
}

func findNewLocation(value int, gardenMaps []GardenMap) int {
	for _, m := range gardenMaps {
		if value >= m.from && value < (m.from+m.rng) {
			return m.to + (value - m.from)
		}
	}
	return value
}
