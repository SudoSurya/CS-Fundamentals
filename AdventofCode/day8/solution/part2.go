package main

import (
	"strings"
)

func PART2() {

	contents := readInput()
	nodeMap := map[string]Node{}
	endWithA := []string{}
	instructions := contents[0]
	for _, line := range strings.Split(contents[1], "\n") {
		lineSplit := strings.Split(line, " =")
		if strings.HasSuffix(lineSplit[0], "A") {
			endWithA = append(endWithA, lineSplit[0])
		}
		nodeSplits := strings.Split(lineSplit[1], ",")
		R := strings.Split(nodeSplits[0], "(")
		L := strings.Split(nodeSplits[1], ")")
		nodeMap[lineSplit[0]] = Node{strings.Trim(R[1], " "), strings.Trim(L[0], " ")}
	}
	var finalInstructions string
	for i := 0; i < 100; i++ {
		finalInstructions += instructions
	}
    pathCounts := []int{}
	for _, ele := range endWithA {
		dest := ele
		count := 0
		for !strings.HasSuffix(dest, "Z") {
			target := finalInstructions[count]
			if string(target) == "R" {
				dest = nodeMap[dest].R
			} else {
				dest = nodeMap[dest].L
			}
			count++
		}
        pathCounts = append(pathCounts, count)
	}
    LCM(pathCounts)
}
