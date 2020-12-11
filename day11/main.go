package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	FLOOR, EMPTY_SEAT, OCCUPIED_SEAT = ".", "L", "#"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	checkErr(err)

	lines := strings.Split(string(content), "\n")
	seats := make([][]string, len(lines))
	seatspart2 := make([][]string, len(lines))
	for i, line := range lines {
		seats[i] = make([]string, len(line))
		seatspart2[i] = make([]string, len(line))
		for j:=0;j<len(line);j++ {
			seats[i][j] = string(line[j])
			seatspart2[i][j] = string(line[j])
		}
	}

	var numOfChanges int = 1
	for numOfChanges > 0 {
		seats, numOfChanges = moveOfSeatsPlease(seats, 4, false)
	}

	fmt.Println("part 1:",countState(seats, OCCUPIED_SEAT))

	numOfChanges = 1
	for numOfChanges > 0 {
		seatspart2, numOfChanges = moveOfSeatsPlease(seatspart2, 5, true)
	}

	fmt.Println("part 2:", countState(seatspart2, OCCUPIED_SEAT))
}

func countState(s [][]string, desiredState string) (count int) {
	for x:=0;x<len(s);x++ {
		for y:=0;y<len(s[x]);y++ {
			if s[x][y] == desiredState {
				count++
			}
		}
	}
	return count
}

func moveOfSeatsPlease(s [][]string, maxAdjacents int, checkViews bool) ([][]string, int) {
	copiedS := make([][]string, len(s))
	copy(copiedS, s)
	numOfChanges := 0
	for x:=0;x<len(s);x++ {
		copiedS[x] = make([]string, len(s[x]))
		copy(copiedS[x], s[x])
		for y:=0;y<len(s[x]);y++ {
			switch s[x][y] {
			case EMPTY_SEAT:
				// If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
				if occuppiedAdjacents(s, x, y, checkViews) == 0 {
					copiedS[x][y] = OCCUPIED_SEAT
					numOfChanges++
				}
			case OCCUPIED_SEAT:
				// If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
				if occuppiedAdjacents(s, x, y, checkViews) >= maxAdjacents {
					copiedS[x][y] = EMPTY_SEAT
					numOfChanges++
				}
			}
		}
	}

	return copiedS, numOfChanges
}

func occuppiedAdjacents(s [][]string, x int, y int, checkViews bool) (adjacents int) {
	for i:=x-1;i<=x+1;i++ {
		if i<0 || i>= len(s) {
			continue
		}

		for j:=y-1;j<=y+1;j++ {
			if j<0 || j>= len(s[i]) || (j == y && i == x) {
				continue
			}

			if s[i][j] == OCCUPIED_SEAT {
				adjacents++
			} else if s[i][j] == FLOOR && checkViews {
				tmpi, tmpj := i, j

				for {
					tmpi = incDecOrUnmodVar(tmpi, i, x)
					tmpj = incDecOrUnmodVar(tmpj, j, y)

					if tmpi < 0 || tmpj < 0 || tmpi >= len(s) || tmpj >= len(s[i]) {
						break
					}

					if val := s[tmpi][tmpj]; val != FLOOR {
						if val == OCCUPIED_SEAT {
							adjacents++
						}

						break
					}
				}
			}
		}
	}

	return adjacents
}

func incDecOrUnmodVar(tmpi int, i int, x int) int {
	if i<x {
		tmpi--
	} else if i>x {
		tmpi++
	}

	return tmpi
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}