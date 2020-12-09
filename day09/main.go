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

	numbers := make([]int, len(lines))
	for i, v := range lines {
		numbers[i], _ = strconv.Atoi(v)
	}

	resultp1 := part1(numbers)
	fmt.Println("part 1 result", resultp1)

	resultp2 := part2(numbers)
	fmt.Println("part 2 result", resultp2)
}

func part1(numbers []int) int {
	for i:=25;i<len(numbers);i++ {
		var isValid bool
		for j:=i-25;j<i;j++ {
			for k:=j+1;k<i;k++ {
				if isValid = numbers[j] + numbers[k] == numbers[i]; isValid {
					j, k = len(numbers), len(numbers)
				}
			}
		}

		if !isValid {
			return numbers[i]
		}
	}

	return 0
}

func part2(numbers []int) int {
	part1result := part1(numbers)

	for i := range numbers {
		for j:=i+1;j<len(numbers);j++ {
			sum := 0
			for k:=i;k<j+1;k++ {
				// that's our window. From i to j+1
				sum = sum + numbers[k]
			}
			if sum == part1result {
				return Min(numbers[i:j+1]) + Max(numbers[i:j+1])
			}
		}
	}

	return 0
}

func Min(in []int) (min int) {
	min = in[0]
	for _, v := range in {
		if v < min {
			min = v
		}
	}
	return min
}

func Max(in []int) (max int) {
	max = in[0]
	for _, v := range in {
		if v > max {
			max = v
		}
	}
	return max
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}