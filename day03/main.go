package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const TREE_VALUE = "#"

func main() {
	content, err := ioutil.ReadFile("input.txt")
	checkErr(err)

	lines := strings.Split(string(content), "\n")
	grid2d := make([][]string, len(lines))

	for y, v := range lines {
		grid2d[y] = make([]string, len(v))
		for x :=0; x < len(v); x++ {
			grid2d[y][x] = string(v[x])
		}
	}

	part1 := countTrees(grid2d, 3, 1)
	fmt.Println("part1", part1)

	part2 := plusIt(countTrees, grid2d, []*XY{{3, 1},{1, 1}, {5, 1}, {7, 1}, {1, 2}})
	fmt.Println("part2", part2)

}

func plusIt(trees func(grid2d [][]string, stepsX int, stepsY int) (trees int), grid2d [][]string, xies []*XY) (plus int) {
	plus = 1
	for _, v := range xies {
		plus = plus * trees(grid2d, v.X, v.Y)
	}
	return plus
}

func countTrees(grid2d [][]string, stepsX int, stepsY int) (trees int) {
	var x, y int
	for y < len(grid2d) {
		x = x % len(grid2d[y])
		trees, x, y = trees + strings.Count(grid2d[y][x], TREE_VALUE), x+stepsX, y+stepsY
	}
	return trees
}

type XY struct {
	X int
	Y int
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}