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

	part1res, part2res := 0, 0

	for _, line := range strings.Split(string(content), "\n") {
		part1res += execute(line, solveOperationLogicPriority)
		part2res += execute(line, solveOperationMultiplicationPriority)
	}

	fmt.Println("part1res", part1res)
	fmt.Println("part2res", part2res)
}

func execute(line string, solver func (input string) string) int {
	line = strings.ReplaceAll(line, " ", "")

	for strings.Contains(line, "(") {
		splittedLeft := strings.Split(line, "(")
		for i:=0;i<len(splittedLeft);i++ {
			if strings.Contains(splittedLeft[i], ")") {
				splittedRight := strings.Split(splittedLeft[i], ")")

				line = strings.ReplaceAll(line, "(" + splittedRight[0]  + ")", solver(splittedRight[0]))
			}
		}
	}

	sum, _ := strconv.Atoi(solver(line))

	return sum
}

func solveOperationLogicPriority(input string) string {
	splitted := strings.Split(strings.ReplaceAll(input, "+", "*"), "*")
	nums := make([]int, len(splitted))
	for i, value := range splitted {
		nums[i], _ = strconv.Atoi(value)
	}

	numPos := 1
	accum := nums[0]

	for i:=0;i<len(input);i++ {
		switch input[i : i+1] {
		case "+":
			accum += nums[numPos]
			numPos++
		case "*":
			accum *= nums[numPos]
			numPos++
		}
	}

	return strconv.Itoa(accum)
}

func solveOperationMultiplicationPriority(input string) string {
	for strings.Contains(input, "+") {
		for _, v := range strings.Split(input, "*") {

			if splittedSum := strings.Split(v, "+"); len(splittedSum) > 1 {
				num1, _ := strconv.Atoi(splittedSum[0])
				num2, _ := strconv.Atoi(splittedSum[1])

				input = strings.ReplaceAll(input, v, strings.ReplaceAll(v, splittedSum[0] + "+" + splittedSum[1], strconv.Itoa(num1+num2)))

				break
			}
		}
	}

	return solveOperationLogicPriority(input)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}