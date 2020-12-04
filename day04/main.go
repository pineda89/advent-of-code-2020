package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	checkErr(err)

	lines := strings.Split(string(content), "\n")

	countValids1, countValids2 := 0, 0

	passportBuff := bytes.Buffer{}
	for _, line := range lines {
		if len(line) == 0 {
			passport := &Passport{}
			passport.rawTextToStruct(passportBuff.String())

			if passport.isValidPart1() {
				countValids1++
			}
			if passport.isValidPart2() {
				countValids2++
			}

			passportBuff = bytes.Buffer{}
		} else {
			passportBuff.WriteString(line)
			passportBuff.WriteString("\n")
		}
	}

	fmt.Println("countValids1", countValids1)
	fmt.Println("countValids2", countValids2)

}

type Passport struct {
	ecl string
	pid string
	eyr string
	hcl string
	byr string
	iyr string
	cid string
	hgt string
}

func (passport *Passport) rawTextToStruct(input string) {
	lines := strings.Split(strings.Replace(input, " ", "\n", -1), "\n")
	data := make(map[string]string)
	for _, line := range lines {
		kv := strings.Split(line, ":")
		if len(kv) > 1 {
			data[kv[0]] = kv[1]
		}
	}

	passport.ecl, passport.pid, passport.eyr, passport.hcl, passport.byr, passport.iyr, passport.cid, passport.hgt = data["ecl"], data["pid"], data["eyr"], data["hcl"], data["byr"], data["iyr"], data["cid"], data["hgt"]
}

func (passport *Passport) isValidPart1() bool {
	return len(passport.byr) > 0 && len(passport.iyr) > 0 && len(passport.eyr) > 0 && len(passport.hgt) > 0 && len(passport.hcl) > 0 && len(passport.ecl) > 0 && len(passport.pid) > 0
}

func (passport *Passport) isValidPart2() bool {
	if !passport.isValidPart1() {
		return false
	}

	byr, _ := strconv.Atoi(passport.byr)
	iyr, _ := strconv.Atoi(passport.iyr)
	eyr, _ := strconv.Atoi(passport.eyr)
	hgt, _ := strconv.Atoi(passport.hgt[:len(passport.hgt)-2])

	validHgt := false
	if passport.hgt[len(passport.hgt)-2:] == "cm" {
		validHgt = hgt >= 150 && hgt <= 193
	} else if passport.hgt[len(passport.hgt)-2:] == "in" {
		validHgt = hgt >= 59 && hgt <= 76
	}

	hclMatched, _ := regexp.MatchString("#[0-9a-f]", passport.hcl)
	pidMatched, _ := regexp.MatchString("[0-9]", passport.pid)

	validEcl := false
	switch passport.ecl {
		case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
			validEcl = true
	}

	return (byr >= 1920 && byr <= 2002) && (iyr >= 2010 && iyr <= 2020) && (eyr >= 2020 && eyr <= 2030) && validHgt && (hclMatched && len(passport.hcl) == 7) && (pidMatched && len(passport.pid) == 9) && validEcl
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}