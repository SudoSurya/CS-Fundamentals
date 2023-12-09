package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var numberRe = regexp.MustCompile(`-?\d+(\.\d+)?`)

func main() {
	input := readInput()
	reports := [][]int{}
	for _, line := range input {
		nums := numberRe.FindAllString(line, -1)
		reports = append(reports, strToInt(nums))
	}
	extrapolatedVals := []int{}
    reportSeq := [][]int{}
	for _, report := range reports {
		last := 0
        begins := []int{}
		for !allZero(report) {
			last += report[len(report)-1]
            begins = append(begins, report[0])
			report = makeDiffArray(report)
		}
		extrapolatedVals = append(extrapolatedVals, last)
        reportSeq = append(reportSeq, begins)
	}
	fmt.Println(sums(extrapolatedVals))

	exPolatedVals := 0
	for _, seq := range reportSeq {
		exPolatedVals += getLeftSeq(seq)
	}
	fmt.Println(exPolatedVals)
}
func getLeftSeq(nums []int) int {
	overlap := 0
	for i := len(nums) - 1; i >= 0; i-- {
		overlap = nums[i] - overlap
	}
	return overlap
}
func sums(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func strToInt(str []string) []int {
	nums := make([]int, len(str))
	for i, s := range str {
		nums[i], _ = strconv.Atoi(s)
	}
	return nums
}
func allZero(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}
func makeDiffArray(nums []int) []int {
	diff := make([]int, len(nums)-1)
	for i := 1; i < len(nums); i++ {
		diff[i-1] = nums[i] - nums[i-1]
	}
	return diff
}
func readInput() []string {
	inputByte, err := os.ReadFile("../input/prod")
	if err != nil {
		panic(err)
	}
	inputByte = inputByte[:len(inputByte)-1]
	inputStr := string(inputByte)
	input := strings.Split(inputStr, "\n")
	return input
}
