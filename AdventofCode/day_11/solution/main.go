package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	contents := readInput()
	horizontalHashIdx := []int{}
	hashIdx := [][]int{}
	verticalHashIdx := make(map[int]bool, len(contents[0]))
	for idx, line := range contents {
		if strings.Contains(line, "#") {
			for i := 0; i < len(line); i++ {
				if string(line[i]) == "#" {

					hashIdx = append(hashIdx, []int{idx, i})
					verticalHashIdx[i] = true
				}
			}
			continue
		}
		horizontalHashIdx = append(horizontalHashIdx, idx)
	}
	vetIdx := []int{}
	for i := 0; i < len(contents[0]); i++ {
		if !verticalHashIdx[i] {
			vetIdx = append(vetIdx, i)
		}
	}

	total := 0
	for i := 0; i < len(hashIdx); i++ {
		r1, c1 := hashIdx[i][0], hashIdx[i][1]
		for j := i + 1; j < len(hashIdx); j++ {
			r2, c2 := hashIdx[j][0], hashIdx[j][1]
			for k := min(r1, r2); k < max(r1, r2); k++ {
				if needle(horizontalHashIdx, k) {
					total += 1000000
				} else {
					total += 1
				}
			}
			for k := min(c1, c2); k < max(c1, c2); k++ {
				if needle(vetIdx, k) {
					total += 1000000
				} else {
					total += 1
				}
			}
		}
	}
	fmt.Println("total",total)
}
func needle(haystack []int, needle int) bool {
	for _, val := range haystack {
		if val == needle {
			return true
		}
	}
	return false
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
