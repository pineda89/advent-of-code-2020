package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	checkErr(err)

	seats := make([]*Seat, 0)
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		seat := &Seat{}
		seat.lineToObj(line)
		seats = append(seats, seat)
	}

	sort.Sort(SortableSeat(seats))
	var firstAnswer, secondAnswer = seats[0].seatId, 0
	for i:=1;i<len(seats);i++ {
		firstAnswer = int(math.Max(float64(firstAnswer), float64(seats[i].seatId)))

		if seats[i].seatId - seats[i-1].seatId != 1 {
			// check if is correlative
			secondAnswer = seats[i-1].seatId + 1
		}
	}

	fmt.Println("firstAnswer", firstAnswer)
	fmt.Println("secondAnswer", secondAnswer)
}

// findRow and findColumn works equals. Replacing F and L by 0, B and R by 1, and parsing the string to binary number
func findRow(input string) int {
	tmp := strings.ReplaceAll(strings.ReplaceAll(input[:len(input)-3], "F", "0"), "B", "1")
	rowId, _ := strconv.ParseInt(tmp, 2, 64)
	return int(rowId)
}

func findColumn(input string) int {
	tmp := strings.ReplaceAll(strings.ReplaceAll(input[len(input)-3:], "L", "0"), "R", "1")
	colId, _ := strconv.ParseInt(tmp, 2, 64)
	return int(colId)
}

type Seat struct {
	row    int
	column int
	seatId int
}

func (seat *Seat) lineToObj(input string) {
	seat.row = findRow(input)
	seat.column = findColumn(input)
	seat.seatId = seat.row * 8 + seat.column
}

type SortableSeat []*Seat

func (a SortableSeat) Len() int           { return len(a) }
func (a SortableSeat) Less(i, j int) bool {
	if a[i].row == a[j].row {
		if a[i].column == a[j].column {
			// same row and same column. Sort by seatId
			return a[i].seatId < a[j].seatId
		} else {
			// same row, different column. Sort by column
			return a[i].column < a[j].column
		}
	} else {
		// different row. Sort by row
		return a[i].row < a[j].row
	}
}
func (a SortableSeat) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}