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

	val := calculateMadafaka(lines, 2020, 1, 0, 0, 1)
	fmt.Println(val)
	val = calculateMadafaka(lines, 2020, 2, 0, 0, 1)
	fmt.Println(val)
}

func calculateMadafaka(lines []string, expectedValue int, stepsLeft int, currentIndex, currentSum int, currentPlus int) int {
	for i:=currentIndex;i<len(lines);i++ {
		firstNum := parseNum(lines[i])
		if stepsLeft > 0 {
			if tmpValue := calculateMadafaka(lines, expectedValue, stepsLeft-1, i, currentSum + firstNum, currentPlus * firstNum); tmpValue != 0 {
				return tmpValue
			}
		} else {
			if currentSum + firstNum == expectedValue {
				return currentPlus * firstNum
			}
		}
	}
	return 0
}

func parseNum(input string) int {
	num, _ := strconv.Atoi(strings.Split(input, "\r")[0])
	return num
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
