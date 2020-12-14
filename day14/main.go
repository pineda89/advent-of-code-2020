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

	part1result := part1(lines)
	fmt.Println("part1result", part1result)

	part2result := part2(lines)
	fmt.Println("part2result", part2result)
}

func part1(lines []string) (sum int) {
	memory := make(map[uint64]uint64)
	var currentMask string
	for _, line := range lines {
		if lineFields := strings.Split(line, " = "); lineFields[0] == "mask" {
			currentMask = lineFields[1]
		} else {
			address, _ := strconv.ParseUint(strings.Split(strings.Split(lineFields[0],"[")[1], "]")[0], 10, 64)
			value, _ := strconv.ParseUint(lineFields[1], 10, 64)

			memory[address] = applyMask(currentMask, value)
		}
	}

	for _, v := range memory {
		sum += int(v)
	}
	return sum
}

func applyMask(mask string, value uint64) uint64 {
	andMask, _ := strconv.ParseUint(strings.ReplaceAll(mask, "X", "1"), 2, 64)
	orMask, _ := strconv.ParseUint(strings.ReplaceAll(mask, "X", "0"), 2, 64)

	return (value & andMask) | orMask
}

func part2(lines []string) (sum int) {
	currentMask := ""
	memory := make(map[uint64]uint64)

	for _, line := range lines {
		if splittedV := strings.Split(line, " = "); splittedV[0] == "mask" {
			currentMask = splittedV[1]
		} else {
			address, _ := strconv.Atoi(strings.Split(strings.Split(line,"[")[1], "]")[0])
			value, _ := strconv.ParseUint(strings.Split(line, " = ")[1], 10, 64)

			permutedsMasks := findAllPosiblesValues("", currentMask, fmt.Sprintf("%036b", address))
			for _, generatedAddr := range permutedsMasks {
				memory[generatedAddr] = value
			}
		}
	}

	for _, v := range memory {
		sum += int(v)
	}
	return sum
}


func findAllPosiblesValues(currentMask, maskPending, address36bitsformat string) []uint64 {
	if len(maskPending) == 0 {
		val, _ := strconv.ParseUint(currentMask, 2, 64)
		return []uint64{val}
	}

	checkingValue := maskPending[0:1]
	maskPending = maskPending[1:]

	switch checkingValue {
	case "0":
		// if is 0, we append the same value as original address to the current mask
		originalValue := address36bitsformat[len(currentMask) : len(currentMask)+1]
		return findAllPosiblesValues(currentMask + originalValue, maskPending, address36bitsformat)
	case "1":
		// if is 1, we append 1 to the current mask
		return findAllPosiblesValues(currentMask + "1", maskPending, address36bitsformat)
	case "X":
		// if is X, we check the 2 options (0 and 1)
		f := findAllPosiblesValues(currentMask + "0", maskPending, address36bitsformat)
		f2 := findAllPosiblesValues(currentMask + "1", maskPending, address36bitsformat)
		return append(f, f2...)
	}

	return []uint64{}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}