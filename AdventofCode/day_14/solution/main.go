package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	contents := readInput()
	contents = rotate90(contents)
    zeros := make([]int, len(contents[0]))
	for _, val := range contents {
		val = reverseStr(val)
		splitVal := strings.Split(val, "#")
        for i := 0; i < len(splitVal); i++ {
            splitVal[i] = shiftZerosToLeft(splitVal[i])
        }
        val = strings.Join(splitVal, "#")
        for i := 0; i < len(val); i++ {
            if val[i] == 'O' {
                zeros[i]++
            }
        }
	}
    for i:= 0; i < 1000000000; i++ {
        fmt.Println(i)
    }
    load := 0
    weight := len(zeros)
    for _, val := range zeros {
        if val > 0 {
            load += (val * weight)
        }
        weight--
    }
    fmt.Println(load)
}
func shiftZerosToLeft(s string) string {
	runes := []rune(s)

	// Iterate through the string and shift '0's to the left
	writeIndex := 0
	for _, char := range runes {
		if char == 'O' {
			runes[writeIndex] = 'O'
			writeIndex++
		}
	}

	// Fill the remaining positions with other characters
	for i := writeIndex; i < len(runes); i++ {
		runes[i] = '.'
	}

	return string(runes)
}
func swap(val string, R, L int) string {
	// Convert string to a rune slice
	runes := []rune(val)

	// Swap elements in the rune slice
	runes[R], runes[L] = runes[L], runes[R]

	// Convert the rune slice back to a string
	return string(runes)
}
func reverseStr(str string) string {
	var result string
	for i := len(str) - 1; i >= 0; i-- {
		result += string(str[i])
	}
	return result
}

func rotate90Left(matrix []string) []string {
    result := make([]string, len(matrix[0]))
    for i := 0; i < len(matrix[0]); i++ {
        var row string
        for j := 0; j < len(matrix); j++ {
            row += string(matrix[j][len(matrix[0])-1-i])
        }
        result[i] = row
    }
    return result
}
func rotate90(matrix []string) []string {
	result := make([]string, len(matrix[0]))
	for i := 0; i < len(matrix[0]); i++ {
		var row string
		for j := len(matrix) - 1; j >= 0; j-- {
			row += string(matrix[j][i])
		}
		result[i] = row
	}
	return result
}


func readInput() []string {
	bytes, err := os.ReadFile("../input/prod")
	if err != nil {
		panic(err)
	}
	bytes = bytes[:len(bytes)-1]
	str := strings.Split(string(bytes), "\n")
	return str
}
