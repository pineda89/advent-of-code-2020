package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	checkErr(err)

	lines := strings.Split(string(content), "\n")

	movements := make([]*Movement, len(lines))

	for i, line := range lines {
		val, _ := strconv.ParseFloat(line[1:], 64)
		movements[i] = &Movement{action: line[:1], quantity: float32(val)}
	}

	part1result := part1(movements)
	fmt.Println("part1result", part1result)

	part2result := part2(movements)
	fmt.Println("part2result", part2result)

	// both can be implemented as part2, but I'm lazy, it's saturday!
}

func part1(movements []*Movement) float32 {
	positions := make(map[string]float32)
	directions := []string{"E", "S", "W", "N"}

	currentDirection := 0 // initial direction
	for _, movement := range movements {
		switch movement.action {
		case "F":
			positions[directions[currentDirection]] += movement.quantity
		case "N", "S", "E", "W":
			positions[movement.action] += movement.quantity
		case "L":
			currentDirection = (currentDirection - (int(movement.quantity)/90)) % 4
			if currentDirection < 0 {
				currentDirection = len(directions) + currentDirection
			}
		case "R":
			currentDirection = (currentDirection + (int(movement.quantity)/90)) % 4
		}
	}

	return float32(math.Abs(float64(positions["S"] - positions["N"])) + math.Abs(float64(positions["E"] - positions["W"])))
}

func part2(movements []*Movement) float32 {
	directions := map[string]Vector2{
		"E": {X: 1, Y: 0},
		"W": {X: -1, Y: 0},
		"N": {X: 0, Y: 1},
		"S": {X: 0, Y: -1},
	}

	waypointCoordinates, shipCoordinates := Vector2{10, 1}, Vector2{}

	for _, movement := range movements {
		switch movement.action {
		case "F":
			shipCoordinates = shipCoordinates.Add(waypointCoordinates.Mul(movement.quantity))
		case "N", "S", "E", "W":
			waypointCoordinates = waypointCoordinates.Add(directions[movement.action].Mul(movement.quantity))
		case "L":
			for i:=0;i<int(movement.quantity/90);i++ {
				waypointCoordinates = Vector2{waypointCoordinates.Y * -1, waypointCoordinates.X}
			}
		case "R":
			for i:=0;i<int(movement.quantity/90);i++ {
				waypointCoordinates = Vector2{waypointCoordinates.Y, waypointCoordinates.X * -1}
			}
		}
	}

	return float32(math.Abs(float64(shipCoordinates.X)) + math.Abs(float64(shipCoordinates.Y)))
}

type Movement struct {
	action string
	quantity float32
}

type Vector2 struct {
	X float32
	Y float32
}

func (v Vector2) Add(other Vector2) Vector2 {
	return Vector2{X: v.X + other.X, Y: v.Y + other.Y}
}

func (v Vector2) Mul(i float32) Vector2 {
	return Vector2{X: v.X * i, Y: v.Y * i}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}