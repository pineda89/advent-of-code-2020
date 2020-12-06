package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	checkErr(err)

	lines := strings.Split(string(content), "\n")

	groups := make([]*Group, 0)
	currentGroup := &Group{answers: make(map[string]int)}
	for _, line := range lines {
		if len(line) == 0 {
			groups = append(groups, currentGroup)
			currentGroup = &Group{answers: make(map[string]int)}
		} else {
			person := &Person{answers: make(map[string]int)}
			person.rawTextToStruct(line)
			for k, v := range person.answers {
				currentGroup.answers[k] += v
			}
			currentGroup.persons = append(currentGroup.persons, person)
		}
	}

	firstAnswer, secondAnswer := 0, 0
	for _, group := range groups {
		firstAnswer += len(group.answers)

		for question := range group.answers {
			allAnswers := true
			for _, person := range group.persons {
				if person.answers[question] == 0 {
					allAnswers = false
				}
			}

			if allAnswers {
				secondAnswer++
			}
		}
	}

	fmt.Println("firstAnswer", firstAnswer)
	fmt.Println("secondAnswer", secondAnswer)
}

type Group struct {
	persons []*Person
	answers map[string]int
}

type Person struct {
	answers map[string]int
}

func (person *Person) rawTextToStruct(input string) {
	for _, v := range input {
		person.answers[string(v)]++
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}