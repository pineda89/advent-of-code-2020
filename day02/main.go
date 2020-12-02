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

	counterValidsMethod1 := 0
	counterValidsMethod2 := 0

	for _, v := range lines {
		pw := &PasswordValidator{}
		pw.LineToPasswordValidator(v)
		if pw.isValidMethodOne() {
			counterValidsMethod1++
		}
		if pw.isValidMethodTwo() {
			counterValidsMethod2++
		}
	}

	fmt.Println("counterValidsMethod1", counterValidsMethod1)
	fmt.Println("counterValidsMethod2", counterValidsMethod2)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type PasswordValidator struct {
	MinTimes int
	MaxTimes int
	Char string
	Chain string
}

func (pw *PasswordValidator) LineToPasswordValidator(input string) {
	pw.MinTimes, _ = strconv.Atoi(strings.Split(input, "-")[0])
	pw.MaxTimes, _ = strconv.Atoi(strings.Split(strings.Split(input, "-")[1], " ")[0])
	pw.Char = strings.Split(strings.Split(input, " ")[1], ":")[0]
	pw.Chain = strings.Split(input, ": ")[1]
}

func (pw *PasswordValidator) isValidMethodOne() bool {
	count := strings.Count(pw.Chain, pw.Char)
	return count >= pw.MinTimes && count <= pw.MaxTimes
}

func (pw *PasswordValidator) isValidMethodTwo() bool {
	return (pw.Chain[pw.MinTimes-1:pw.MinTimes] == pw.Char && pw.Chain[pw.MaxTimes-1:pw.MaxTimes] != pw.Char) ||
		(pw.Chain[pw.MinTimes-1:pw.MinTimes] != pw.Char && pw.Chain[pw.MaxTimes-1:pw.MaxTimes] == pw.Char)
}