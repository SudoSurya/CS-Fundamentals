package main

import (
	"fmt"
	"os"
	"strings"
)

type Node struct {
	L string
	R string
}

func main() {
	contents := readInput()
	nodeMap := map[string]Node{}
	instructions := contents[0]
	for _, line := range strings.Split(contents[1], "\n") {
		lineSplit := strings.Split(line, " =")
		nodeSplits := strings.Split(lineSplit[1], ",")
		R := strings.Split(nodeSplits[0], "(")
		L := strings.Split(nodeSplits[1], ")")
		nodeMap[lineSplit[0]] = Node{strings.Trim(R[1], " "), strings.Trim(L[0], " ")}
	}
	var finalInstructions string
	for i := 0; i < 100; i++ {
		finalInstructions += instructions
	}
	dest := "AAA"
	count := 0
	for dest != "ZZZ" {
		target := finalInstructions[count]
		if string(target) == "R" {
			dest = nodeMap[dest].R
		} else {
			dest = nodeMap[dest].L
		}
		count++
	}
	fmt.Println("part1",count)
    fmt.Println("PART2")
    PART2()
}

func readInput() []string {
	inputByte, err := os.ReadFile("../input/prod")

	if err != nil {
		panic(err)
	}
	inputByte = inputByte[:len(inputByte)-1]
	inputStr := string(inputByte)
	input := strings.Split(inputStr, "\n\n")
	return input
}
