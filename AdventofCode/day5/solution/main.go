package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	numberRe = regexp.MustCompile(`\d+`)
)

func main() {
	fmt.Println("start")
    part1()
    part2()
}

func convertToAlmanc(seedItems *[]int, mapItems *[]string) {
	for item := 0; item < len(*seedItems); item++ {
		for i := 1; i < len(*mapItems); i++ {
			mapNums := numberRe.FindAllString((*mapItems)[i], -1)
			seedItemCopy := (*seedItems)[item]
			seedPos, foundSeed := getSeedNumber(seedItemCopy, ints(mapNums[0]), ints(mapNums[1]), ints(mapNums[2]))
			if foundSeed {
				(*seedItems)[item] = seedPos
				break
			}
		}
	}
}

func getSeedNumber(seedNumber int, destinationRange, sourceRange, rangeLength int) (int, bool) {
	if seedNumber >= sourceRange && seedNumber <= (sourceRange+rangeLength) {
		diff := seedNumber - sourceRange
		return destinationRange + diff, true
	}
	return seedNumber, false
}

func ints(str string) int {
	strInt, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return strInt
}
func findMinElement(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	// Initialize min with the first element of the array
	min := arr[0]

	// Iterate through the array starting from the second element
	for _, value := range arr[1:] {
		if value < min {
			min = value
		}
	}

	return min
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
