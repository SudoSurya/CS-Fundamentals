package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	contents := readInput()
	expandedGrid := [][]string{}
    horizontalHashIdx := []int{}
	hashIdx := [][]int{}
	for idx, line := range contents {
		expandedGrid = append(expandedGrid, strings.Split(line, ""))
		if strings.Contains(line, "#") {
			continue
		}
        horizontalHashIdx = append(horizontalHashIdx, idx)
		expandedGrid = append(expandedGrid, strings.Split(line, ""))
	}

	grid := make([][]int, len(expandedGrid))
	for i := range grid {
		grid[i] = make([]int, len(expandedGrid[0]))
	}
	for idx, line := range expandedGrid {
		for i, val := range line {
			if strings.Contains(val, "#") {
				grid[idx][i] = 1
			} else {
				grid[idx][i] = 0
			}
		}
	}
	grid = transposeArray(grid)
	newGrid := [][]int{}
	for i := 0; i < len(grid); i++ {
		isWholeRowZero := true
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != 0 {
				isWholeRowZero = false
				break
			}
		}
		if !isWholeRowZero {
			newGrid = append(newGrid, grid[i])
		} else {
			newGrid = append(newGrid, grid[i])
			newGrid = append(newGrid, grid[i])
		}
	}
	newGrid = transposeBack(newGrid)
	for idx, row := range newGrid {
		for i, val := range row {
            if val == 1 {
                hashIdx = append(hashIdx, []int{idx, i})
            }
		}
	}
    total := 0
    for i := 0; i < len(hashIdx); i++ {
        r1, c1 := hashIdx[i][0], hashIdx[i][1]
        for j := i+1; j < len(hashIdx); j++ {
            r2, c2 := hashIdx[j][0], hashIdx[j][1]
            for k := min(r1, r2); k < max(r1, r2); k++ {
                total += 1
            }
            for k := min(c1, c2); k < max(c1, c2); k++ {
                total += 1
            }
        }
    }
    fmt.Println(total)
}
func transposeBack(arr [][]int) [][]int {
	rows := len(arr)
	cols := len(arr[0])

	// Create a new array with swapped rows and columns
	originalArr := make([][]int, cols)
	for j := range originalArr {
		originalArr[j] = make([]int, rows)
	}

	// Perform the reverse transposition
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			originalArr[j][i] = arr[i][j]
		}
	}

	return originalArr
}
func transposeArray(arr [][]int) [][]int {
	rows := len(arr)
	cols := len(arr[0])

	// Create a new array with swapped rows and columns
	transposedArr := make([][]int, cols)
	for j := range transposedArr {
		transposedArr[j] = make([]int, rows)
	}

	// Perform the transposition
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposedArr[j][i] = arr[i][j]
		}
	}

	return transposedArr
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
