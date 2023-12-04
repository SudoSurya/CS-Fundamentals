package main

import (
	"fmt"
	"regexp"
)

var (
	starRe    = regexp.MustCompile(`\*`)
)

func Part2(contents []string) {
	sum := 0
	for line := 0; line < len(contents); line++ {
		starMatches := starRe.FindAllStringIndex(contents[line], -1)
		mul := calculateGearRatios(line, starMatches, contents)
		sum += mul
	}
	fmt.Println("Sum of all gear ratios:", sum)
}

func calculateGearRatios(line int, starMatches [][]int, contents []string) int {
	mul := 0
	for i := 0; i < len(starMatches); i++ {
		foundNums := findAdjacentPartNumbers(line, starMatches[i], contents)
		if len(foundNums) == 2 {
            temp := foundNums[0] * foundNums[1]
			mul += temp
		}
		fmt.Println("foundNums", foundNums)
	}
	return mul
}

func findAdjacentPartNumbers(line int, starMatch []int, contents []string) []int {
	foundNums := []int{}

    // checking previous line for adjacent part numbers
	if line > 0 {
		foundNums = append(foundNums, findPartNumbers(line-1, starMatch, contents)...)
	}

    // checking next line for adjacent part numbers
	if line < len(contents)-1 {
		foundNums = append(foundNums, findPartNumbers(line+1, starMatch, contents)...)
	}
    // checking current line for adjacent part numbers
    foundNums = append(foundNums, findPartNumbers(line, starMatch, contents)...)

	return foundNums
}

func findPartNumbers(line int, starMatch []int, contents []string) []int {
	foundNums := []int{}
	numberIdxs := numberRe.FindAllStringIndex(contents[line], -1)
	numberMatches := numberRe.FindAllString(contents[line], -1)

	for j := 0; j < len(numberIdxs); j++ {
		for k := numberIdxs[j][0]; k <= numberIdxs[j][1]; k++ {
			if containsint(starMatch, k) {
				fmt.Printf("Found at line %d: %v %v\n", line, numberIdxs[j], numberMatches[j])
				foundNums = append(foundNums, stringToInt(numberMatches[j]))
				break
			}
		}
	}

	return foundNums
}

func containsint(arr []int, str int) bool {
    if str >= arr[0] && str <= arr[1] {
        return true
    }
	return false
}

