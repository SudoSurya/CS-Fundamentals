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
	charRe = regexp.MustCompile(`[^\d.]`)
)
func main() {
	contents := readInput()
	validPartNums := []int{}
	sum := 0
	for line := 0; line < len(contents); line++ {
		macthes := numberRe.FindAllString(contents[line], -1)
		numberIdxs := numberRe.FindAllStringIndex(contents[line], -1)
		checkForMatches(macthes, line, line-1, line+1, numberIdxs, contents, &validPartNums, &sum)
	}
	// fmt.Println("validPartNums", validPartNums)
	fmt.Println("sum", sum)
    Part2(contents)
}

func checkForMatches(
	matches []string,
	line, prevLine, nextLine int,
	numberIdxs [][]int, contents []string,
	validPartNums *[]int, sum *int) {
	for i := 0; i < len(matches); i++ {
		checkMatch(matches[i], numberIdxs[i], line, prevLine, nextLine, contents, validPartNums, sum)
	}
}

func checkMatch(match string, indices []int, line, prevLine, nextLine int, contents []string, validPartNums *[]int, sum *int) {
	if checkInLine(match, indices, contents[line], validPartNums, sum) {
		return
	}
	if prevLine >= 0 {
		checkInLine(match, indices, contents[prevLine], validPartNums, sum)
	}
	if nextLine < len(contents) {
		checkInLine(match, indices, contents[nextLine], validPartNums, sum)
	}
}

func checkInLine(match string, indices []int, lineContents string, validPartNums *[]int, sum *int) bool {
	charIdxs := fmt.Sprintf("%v", charRe.FindAllStringIndex(lineContents, -1))
	charIdsParseArr := numberRe.FindAllString(charIdxs, -1)
	isPartNum := false

	for num := indices[0]; num <= indices[1]; num++ {
		matchStr := strconv.Itoa(num)
		if contains(charIdsParseArr, matchStr) {
			isPartNum = true
			break
		}
	}

	if isPartNum {
		*validPartNums = append(*validPartNums, stringToInt(match))
		*sum += stringToInt(match)
	}

	return isPartNum
}

func contains(arr []string, str string) bool {
	for _, s := range arr {
		if s == str {
			return true
		}
	}
	return false
}

func stringToInt(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}

func readInput() []string {
	inputByte, err := os.ReadFile("../input/01-prod.txt")
	if err != nil {
		panic(err)
	}
	inputByte = inputByte[:len(inputByte)-1]
	inputStr := string(inputByte)
	input := strings.Split(inputStr, "\n")
	return input
}
