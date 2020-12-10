package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	checkErr(err)

	lines := strings.Split(string(content), "\n")
	part1result := part1(lines)
	fmt.Println("part1result", part1result)
	part2result := part2(lines)
	fmt.Println("part2result", part2result)
}

func part1(lines []string) int {
	values := stringSliceToIntSlice(lines)
	sort.Ints(values)

	var results = make(map[int]int)

	for i:=1;i<len(values);i++ {
		results[values[i]-values[i-1]]++
	}

	return (results[1] + 1) * (results[3] + 1)
}

func part2(lines []string) int64 {
	values := stringSliceToIntSlice(lines)
	sort.Ints(values)

	accummulationMap := make(map[int]int64)
	accummulationMap[0]++

	for _, i := range values {
		// for each number, we will check the combinations of latest 3 numbers, also accummulateds
		for j:=1;j<=3;j++ {
			accummulationMap[i] += accummulationMap[i-j]
		}
	}

	return accummulationMap[values[len(values)-1]]
}

func stringSliceToIntSlice(input []string) []int {
	result := make([]int, 0)
	for _, line := range input {
		val, err := strconv.Atoi(line)
		if err == nil {
			result = append(result, val)
		}
	}
	return result
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}