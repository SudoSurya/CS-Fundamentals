package main

import (
	"fmt"
	"strings"
)

func part2() {
	contents := readInput()
	pipes := [][]string{}
	startPos := []int{0, 0}
	for idx, line := range contents {
		pipe := strings.Split(line, "")
		for idx2, char := range pipe {
			if char == "S" {
				startPos = []int{idx, idx2}
			}
		}
		pipes = append(pipes, pipe)
	}

	seen := make([][]bool, len(pipes))
	for i := range seen {
		seen[i] = make([]bool, len(pipes[i]))
	}
	seen[startPos[0]][startPos[1]] = true
	queue := Queue{}
	queue.Enqueue(startPos)
	for !queue.IsEmpty() {
		pipe := queue.Dequeue()
		neigbours := findStartNeigbours(pipes, pipe)
        fmt.Println(neigbours)
        for _, neigbour := range neigbours {
            if !seen[neigbour[0]][neigbour[1]] {
                seen[neigbour[0]][neigbour[1]] = true
                queue.Enqueue(neigbour)
            }
        }
	}
    for _, row := range seen {
        fmt.Println(row)
    }
}
