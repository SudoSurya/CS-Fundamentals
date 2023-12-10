package main

import (
	"fmt"
	"os"
	"strings"
)

type Coord struct {
	Sr, Sc int
}

func main() {

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
	loop := make(map[Coord]bool)
    loop[Coord{startPos[0],startPos[1]}] = true
	queue := Queue{}
	queue.Enqueue(startPos)
	for !queue.IsEmpty() {

		ele := queue.Dequeue()
		r := ele[0]
		c := ele[1]
		ch := pipes[r][c]
		if r > 0 && isIn(ch, []string{"S", "|", "J", "L"}) && isIn(pipes[r-1][c], []string{"|", "7", "F"}) && !loop[Coord{r - 1, c}] {
			loop[Coord{r - 1, c}] = true
			queue.Enqueue([]int{r - 1, c})
		}
		if r < len(pipes)-1 && isIn(ch, []string{"S", "|", "7", "F"}) && isIn(pipes[r+1][c], []string{"|", "J", "L"}) && !loop[Coord{r + 1, c}] {
			loop[Coord{r + 1, c}] = true
			queue.Enqueue([]int{r + 1, c})
		}
		if c > 0 && isIn(ch, []string{"S", "-", "J", "7"}) && isIn(pipes[r][c-1], []string{"-", "L", "F"}) && !loop[Coord{r, c - 1}] {
			loop[Coord{r, c - 1}] = true
			queue.Enqueue([]int{r, c - 1})
		}
		if c < len(pipes[r])-1 && isIn(ch, []string{"S", "-", "L", "F"}) && isIn(pipes[r][c+1], []string{"-", "J", "7"}) && !loop[Coord{r, c + 1}] {
			loop[Coord{r, c + 1}] = true
			queue.Enqueue([]int{r, c + 1})
		}
	}
    fmt.Println(len(loop)/2)
}

func isIn(str string, overlap []string) bool {
	for _, substr := range overlap {
		if strings.Contains(str, substr) {
			return true
		}
	}
	return false
}

func notDot(pipes [][]string, pos []int) bool {
	return pipes[pos[0]][pos[1]] != "."
}
func findStartNeigbours(pipes [][]string, startPos []int) [][]int {
	startPoints := [][]int{}
	if startPos[0] > 0 {
		if notDot(pipes, []int{startPos[0] - 1, startPos[1]}) {
			startPoints = append(startPoints, []int{startPos[0] - 1, startPos[1]})
		}
	}
	if startPos[0] < len(pipes)-1 {
		if notDot(pipes, []int{startPos[0] + 1, startPos[1]}) {
			startPoints = append(startPoints, []int{startPos[0] + 1, startPos[1]})
		}
	}
	if startPos[1] > 0 {
		if notDot(pipes, []int{startPos[0], startPos[1] - 1}) {
			startPoints = append(startPoints, []int{startPos[0], startPos[1] - 1})
		}
	}
	if startPos[1] < len(pipes[0])-1 {
		if notDot(pipes, []int{startPos[0], startPos[1] + 1}) {
			startPoints = append(startPoints, []int{startPos[0], startPos[1] + 1})
		}
	}
	return startPoints
}

func readInput() []string {
	inputByte, err := os.ReadFile("../input/sample")
	if err != nil {
		panic(err)
	}
	inputByte = inputByte[:len(inputByte)-1]
	inputStr := string(inputByte)
	input := strings.Split(inputStr, "\n")
	return input
}
