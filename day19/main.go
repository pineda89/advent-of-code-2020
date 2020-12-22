package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var rules = make(map[string]string)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	checkErr(err)

	finalPartContent := bytes.Buffer{}

	for _, line := range strings.Split(string(content), "\n") {
		if rule := strings.Split(line, ": "); len(rule) > 1 {
			rules[rule[0]] = rule[1]
		} else {
			finalPartContent.WriteString(line)
			finalPartContent.WriteString("\n")
		}
	}

	fmt.Println("part1:", len(regexp.MustCompile("(?m)^"+generateRegexp("0")+"$").FindAllString(finalPartContent.String(), -1)))

	rules["8"] = "\"" + generateRegexp("42") + "+\""
	buf := bytes.Buffer{}
	for i := 1; i < 10; i++ {
		buf.WriteString(fmt.Sprintf("|%s{%d}%s{%d}", generateRegexp("42"), i, generateRegexp("31"), i))
	}
	rules["11"] = `"(?:` + buf.String()[1:] + `)"`

	fmt.Println("part2:", len(regexp.MustCompile("(?m)^"+generateRegexp("0")+"$").FindAllString(finalPartContent.String(), -1)))
}

func generateRegexp(rule string) string {
	var finalRule string
	if strings.Contains(rules[rule], "\"") {
		// rule type is exactly one character
		return strings.ReplaceAll(rules[rule], "\"", "")
	} else {
		for _, s := range strings.Split(rules[rule], "|") {
			if len(finalRule) > 0 {
				finalRule += "|"
			}

			fields := strings.Split(s, " ")
			for _, field := range fields {
				if _, err := strconv.Atoi(field); err == nil {
					finalRule += generateRegexp(field)
				}
			}
		}

		return "(?:" + finalRule + ")"
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}