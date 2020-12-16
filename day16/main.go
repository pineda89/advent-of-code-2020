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
	rules, yourTickets, nearbyTickets := parseFile(lines)

	part1res := part1(rules, yourTickets, nearbyTickets)
	fmt.Println("part1res", part1res)

	part2res := part2(rules, yourTickets, nearbyTickets)
	fmt.Println("part2res", part2res)
}

func parseFile(lines []string) ([]Rule, []int, [][]int) {
	rules := make([]Rule, 0)
	yourTickets := make([]int, 0)
	nearbyTickets := make([][]int, 0)

	lineType := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		if strings.Count(line, "-") == 2 {
			// rule
			splittedLine := strings.Split(line, ": ")
			rangesStr := strings.Split(splittedLine[1], " or ")
			firstRange := strings.Split(rangesStr[0], "-")
			secondRange := strings.Split(rangesStr[1], "-")

			firstMin, _ := strconv.Atoi(firstRange[0])
			firstMax, _ := strconv.Atoi(firstRange[1])
			secondMin, _ := strconv.Atoi(secondRange[0])
			secondMax, _ := strconv.Atoi(secondRange[1])

			rules = append(rules, Rule{ name: splittedLine[0], ruleRanges: []RuleRange{{min: firstMin, max: firstMax}, {min: secondMin, max: secondMax}}})
		} else if strings.HasPrefix(line, "your ticket") {
			lineType = 1
		} else if strings.HasPrefix(line, "nearby tickets") {
			lineType = 2
		} else {
			tmpArray := make([]int, 0)
			splittedLine := strings.Split(line, ",")
			for _, num := range splittedLine {
				ticket, _ := strconv.Atoi(num)
				tmpArray = append(tmpArray, ticket)
			}

			if lineType == 1 {
				yourTickets = append(yourTickets, tmpArray...)
			} else if lineType == 2 {
				nearbyTickets = append(nearbyTickets, tmpArray)
			}
		}
	}

	return rules, yourTickets, nearbyTickets
}

func part1(rules []Rule, yourTickets []int, nearbyTickets [][]int) (sum int) {
	var invalidValues []int

	for _, ticketContent := range nearbyTickets {
		for _, num := range ticketContent {
			ruleOk := false
			for _, rule := range rules {
				if (num >= rule.ruleRanges[0].min && num <= rule.ruleRanges[0].max) || (num >= rule.ruleRanges[1].min && num <= rule.ruleRanges[1].max) {
					ruleOk = true
				}
			}

			if !ruleOk {
				invalidValues = append(invalidValues, num)
			}
		}
	}

	for _, v := range invalidValues {
		sum += v
	}

	return sum
}

func part2(rules []Rule, yourTickets []int, nearbyTickets [][]int) int {
	var correctTickets [][]int

nextTicket:
	for _, ticket := range nearbyTickets {
	nextNumberInTicket:
		for _, ticketDigit := range ticket {
			for _, rule := range rules {
				if (ticketDigit >= rule.ruleRanges[0].min && ticketDigit <= rule.ruleRanges[0].max) ||
					(ticketDigit >= rule.ruleRanges[1].min && ticketDigit <= rule.ruleRanges[1].max) {
					continue nextNumberInTicket
				}
			}

			continue nextTicket // not valid. Ignore this, and check the next ticket
		}

		// if we are here means all digits of ticket are ok
		correctTickets = append(correctTickets, ticket)
	}


	// idea is: departure["departure time"] contains a map with all "yourTickets", then if departures["departure time"][40] is true, means ticket is valid
	departures := make(map[string]map[int]bool)
	for _, rule := range rules {
		departures[rule.name] = make(map[int]bool)

		for i:=0;i<len(yourTickets);i++ {
			validTicket := true
			for j:=0;j<len(correctTickets);j++ {
				if !((correctTickets[j][i] >= rule.ruleRanges[0].min && correctTickets[j][i] <= rule.ruleRanges[0].max) ||
					(correctTickets[j][i] >= rule.ruleRanges[1].min && correctTickets[j][i] <= rule.ruleRanges[1].max)) {
					validTicket = false
				}
			}

			if validTicket {
				// if all numbers of this ticket are valids, add to map
				departures[rule.name][i] = validTicket
			}
		}
	}


	rulesCounter := make(map[string]int)
	for len(departures) > 0 {
		for ruleName := range departures {
			if len(departures[ruleName]) == 1 {
				for ticketDigitPosition := range departures[ruleName] {
					rulesCounter[ruleName] = ticketDigitPosition

					for d := range departures {
						delete(departures[d], rulesCounter[ruleName]) // remove key of all children
					}
				}

				delete(departures, ruleName) // remove parent, like mark as processed
			}
		}
	}

	mult := 1
	for name, ticket := range rulesCounter {
		if strings.Contains(name, "departure") {
			mult = mult * yourTickets[ticket]
		}
	}

	return mult
}

type Rule struct {
	name string
	ruleRanges []RuleRange
}

type RuleRange struct {
	min int
	max int
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
