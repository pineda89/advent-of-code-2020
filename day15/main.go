package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	checkErr(err)

	lines := strings.Split(string(content), "\n")

	part1 := doSomething(lines, 2020)
	fmt.Println("part1", part1)

	part2 := doSomething(lines, 30000000)
	fmt.Println("part2", part2)
}

func doSomething(lines []string, expectedI int) (lastTurnValue int) {
	lastTurnMap := make(map[int]int)

	splittedLine := strings.Split(lines[0], ",")
	for i, field := range splittedLine {
		lastTurnValue, _ = strconv.Atoi(field)
		lastTurnMap[lastTurnValue] = i + 1
	}

	for currentTurn:=len(lastTurnMap);currentTurn<expectedI;currentTurn++ {
		if value := lastTurnMap[lastTurnValue]; value == 0 {
			// is new. Start again from zero
			lastTurnMap[lastTurnValue] = currentTurn
			lastTurnValue = 0
		} else {
			// not new, accumulate it
			lastTurnMap[lastTurnValue], lastTurnValue = currentTurn, currentTurn-value
		}
	}

	return lastTurnValue
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}