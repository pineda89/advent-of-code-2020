package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	checkErr(err)

	lines := strings.Split(string(content), "\n")

	fmt.Println("part1", part1(lines))
	fmt.Println("part2", part2(lines))
}

func part1(lines []string) int {
	timestamp, _ := strconv.Atoi(lines[0])
	splittedLine := strings.Split(lines[1], ",")

	nextBus := math.MaxInt64

	for _, v := range splittedLine {
		if busId, err := strconv.Atoi(v); err == nil {
			// busId-(timestamp%busId) : how much time until bus is here again
			if busId-(timestamp%busId) < nextBus-(timestamp%nextBus) {
				nextBus = busId
			}
		}
	}

	return nextBus * (nextBus - (timestamp%nextBus))
}

func part2(lines []string) int {
	splittedLine := strings.Split(lines[1], ",")

	tmpTime := 0
	multiplier := 1

	for busSeq, v := range splittedLine {
		if busId, err := strconv.Atoi(v); err == nil {
			// we are incrementing by multiplier, which is the multiplication of older buses
			// that confirms, next number is compatible with all checked buses. Is the minium jump needed to find the same pattern

			for (tmpTime+busSeq) % busId != 0 {
				tmpTime = tmpTime + multiplier
			}
			multiplier = multiplier * busId
		}
	}

	return tmpTime
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}