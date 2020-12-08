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

	fmt.Println(part1RunCode(lines))
	fmt.Println(part2FixInfiniteLoop(lines))
}

func part1RunCode(lines []string) (accumulator int, infiniteLoop bool) {
	visitedMap := make(map[int]int)

	currentLine := 0
	for currentLine < len(lines) {
		if _, ok := visitedMap[currentLine]; ok {
			infiniteLoop = true
			break
		}
		visitedMap[currentLine]++

		splittedLine := strings.Split(lines[currentLine], " ")
		q, _ := strconv.Atoi(splittedLine[1])

		switch splittedLine[0] {
		case "acc":
			accumulator = accumulator + q
			currentLine++
		case "jmp":
			currentLine = currentLine + q
		case "nop":
			currentLine++
		}
	}

	return accumulator, infiniteLoop
}

func part2FixInfiniteLoop(lines []string) int {
	modifiedLines := make([]string, len(lines))
	copy(modifiedLines, lines)

	for i := range lines {
		if strings.Contains(modifiedLines[i], "jmp") {
			modifiedLines[i] = strings.ReplaceAll(modifiedLines[i], "jmp", "nop")
		} else if strings.Contains(modifiedLines[i], "nop") {
			modifiedLines[i] = strings.ReplaceAll(modifiedLines[i], "nop", "jmp")
		} else {
			// no changes, not need to check this
			continue
		}

		if accumulator, infLoop := part1RunCode(modifiedLines); !infLoop {
			return accumulator
		}

		// restore line as original
		modifiedLines[i] = lines[i]
	}

	return 0
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}