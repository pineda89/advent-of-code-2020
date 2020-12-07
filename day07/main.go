package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var bagsMap = make(map[string]*bag)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	checkErr(err)

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		split := strings.Split(line, " contain ")
		parent := strings.Split(split[0], " bags")[0]
		for _, tmp := range strings.Split(split[1], ", ") {
			tmp = strings.ReplaceAll(tmp, ".", "")
			tmp = strings.ReplaceAll(tmp, " bags", "")
			tmp = strings.ReplaceAll(tmp, " bag", "")

			splittedChild := strings.SplitN(tmp, " ", 2)
			weight, _ := strconv.Atoi(splittedChild[0])
			child := splittedChild[1]

			addChild(parent, child, weight)
		}
	}

	fmt.Println("first part:", len(bagsMap["shiny gold"].part2GetParents()))
	fmt.Println("second part:", bagsMap["shiny gold"].part1ChildrenWeight())
}

func addChild(parentKey string, childKey string, weight int) {
	parent, child := addNode(parentKey), addNode(childKey)

	edge := weightedEdge{parent: parent, child: child, weight: weight}
	parent.addEdge(&edge, false)
	child.addEdge(&edge, true)
}

func addNode(key string) *bag {
	n := bagsMap[key]
	if n == nil {
		n = &bag{key: key}
		bagsMap[key] = n
	}
	return n
}

type bag struct { // node
	key      string
	parents  []weightedEdge
	children []weightedEdge
}

type weightedEdge struct { // edge
	weight int
	parent *bag
	child  *bag
}

func (n *bag) addEdge(edge *weightedEdge, isParent bool) {
	if isParent {
		n.parents = append(n.parents, *edge)
	} else {
		n.children = append(n.children, *edge)
	}
}

func (n *bag) part1ChildrenWeight() (w int) {
	for _, child := range n.children {
		w += child.weight + (child.weight * child.child.part1ChildrenWeight())
	}

	return w
}

func (n *bag) part2GetParents() []*bag {
	data := make(map[string]*bag)

	for _, parent := range n.parents {
		data[parent.parent.key] = parent.parent
		for _, parent2 := range parent.parent.part2GetParents() {
			data[parent2.key] = parent2
		}
	}

	arr := make([]*bag, len(data))
	i := 0
	for _, v := range data {
		arr[i] = v
		i++
	}

	return arr
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}