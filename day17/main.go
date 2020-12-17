package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

const MAX_DIMENSION = 4
const ACTIVE_VALUE = "#"
const INACTIVE_VALUE = "."

func main() {
	content, err := ioutil.ReadFile("input.txt")
	checkErr(err)

	lines := strings.Split(string(content), "\n")

	dimensionsMap := make(map[Coordinates]string)
	for i, line := range lines {
		for j, value := range line {
			dimensionsMap[Coordinates{Values: [MAX_DIMENSION]int{j, i}}] = string(value)
		}
	}

	part1res := execute(dimensionsMap, 3, 6)
	fmt.Println("part1res", part1res)
	part2res := execute(dimensionsMap, 4, 6)
	fmt.Println("part2res", part2res)
}

func execute(dimensionsMap map[Coordinates]string, dimensions, cycles int) int {
	combinations := removeFullyZeroOrEmpty(findCombinations(dimensions))

	for cyclesLeft:=cycles;cyclesLeft>0;cyclesLeft-- {

		for originalValue := range dimensionsMap {
			for _, combination := range combinations {
				if _, ok := dimensionsMap[originalValue.Sum(combination)]; !ok {
					// if key of neighbors not exists, create it (as inactive)
					dimensionsMap[originalValue.Sum(combination)] = INACTIVE_VALUE
				}
			}
		}

		mutatedCycleDimensionsMap := make(map[Coordinates]string)
		for key, currentValue := range dimensionsMap {
			var neighborsCount int
			for _, combination := range combinations {
				if value, exists := dimensionsMap[key.Sum(combination)]; value == ACTIVE_VALUE && exists {
					neighborsCount++
				}
			}

			if neighborsCount == 3 || (currentValue == ACTIVE_VALUE && neighborsCount == 2) {
				// Only "active" it if currently is active and has 2 neighbors, or if has 3 neighbors ignoring the current status
				// if key is not existing on map, I'm suposing is inactive
				mutatedCycleDimensionsMap[key] = ACTIVE_VALUE
			}
		}
		dimensionsMap = mutatedCycleDimensionsMap
	}

	buf := bytes.Buffer{}
	for _, value := range dimensionsMap {
		buf.WriteString(value)
	}
	return strings.Count(buf.String(), ACTIVE_VALUE)
}

type Coordinates struct {
	Values [MAX_DIMENSION]int
}

func (c Coordinates) Sum(c2 Coordinates) Coordinates {
	for i := range c.Values {
		c.Values[i] = c.Values[i] + c2.Values[i]
	}
	return c
}

func findCombinations(dimension int) []Coordinates {
	if dimension == 0 {
		// return one element as zero value in any dimension
		return []Coordinates{{Values: [MAX_DIMENSION]int{}}}
	}
	coordinates := make([]Coordinates, 0)
	for i:=-1;i<=1;i++ {
		combinations := findCombinations(dimension-1)
		for j:=0;j<len(combinations);j++ {
			combinations[j].Values[dimension-1] = i
			coordinates = append(coordinates, combinations[j])
		}
	}
	return coordinates
}

func removeFullyZeroOrEmpty(combinations []Coordinates) []Coordinates {
	newCombinations := make([]Coordinates, 0)
	for i := range combinations {
		valid := false
		for _, v := range combinations[i].Values {
			if v != 0 {
				valid = true
			}
		}
		if valid {
			newCombinations = append(newCombinations, combinations[i])
		}
	}
	return newCombinations
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}