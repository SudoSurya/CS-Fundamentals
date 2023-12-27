package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	contents := readInput()
	lines := []string{}
	for i := 0; i < len(contents); i++ {
		lines = append(lines, contents[i])
	}
	sum := 0
	for i := 0; i < len(lines); i++ {
		row := findReflextion(strings.Split(lines[i], "\n"))
		col := findReflextion(rotate90(strings.Split(lines[i], "\n")))
		sum += (row * 100)
		sum += col
	}
	fmt.Println(sum)
}
func isRowGreater(row, col int) bool {
	return row > col
}
func findReflextion(lines []string) int {
	for i := 1; i < len(lines); i++ {

		above := reverseStrSlice(lines[:i])
		below := lines[i:]
		if len(above) > len(below) {
			above = above[:len(below)]
		}

		below = below[:len(above)]

        fmt.Println("---- above")
        for _, val := range above {
            fmt.Println(val)
        }
        fmt.Println("---- below")
        for _, val := range below {
            fmt.Println(val)
        }
		if equalSlices(above, below) {
			return i
		}

	}
	return 0
}

func findMax(nums []int) int {
	max := nums[0]
	for _, val := range nums {
		if val > max {
			max = val
		}
	}
	return max
}

func getMid(nums []int) int {
	if len(nums)%2 == 0 {
		return nums[len(nums)/2]
	}
	return nums[len(nums)/2+1]
}
func isContinous(nums []int) bool {
	copyNums := nums
	sort.Ints(copyNums)
	for i := 0; i < len(copyNums)-1; i++ {
		if copyNums[i+1]-copyNums[i] != 1 {
			return false
		}
	}
	return true
}
func isStringsEqual(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	for i, val := range str1 {
		if val != rune(str2[i]) {
			return false
		}
	}
	return true
}

func reverseStrSlice(slice []string) []string {
	reversedSlice := make([]string, len(slice))
	for i := 0; i < len(slice); i++ {
		reversedSlice[i] = slice[len(slice)-i-1]
	}
	return reversedSlice
}
func isStringSliceEqual(slice1 []string, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i, val := range slice1 {
		if val != slice2[i] {
			return false
		}
	}
	return true
}
func transposeArray(arr []string) []string {
	if len(arr) == 0 {
		return nil
	}

	numRows, numCols := len(arr[0]), len(arr)

	newArr := make([]string, numRows)
	for i := 0; i < numRows; i++ {
		newArr[i] = ""
		for j := 0; j < numCols; j++ {
			newArr[i] += string(arr[j][i])
		}
	}
	return newArr
}
func readInput() []string {
	inputByte, err := os.ReadFile("../input/sample")
	if err != nil {
		panic(err)
	}
	inputByte = inputByte[:len(inputByte)-1]
	inputStr := string(inputByte)
	input := strings.Split(inputStr, "\n\n")
	return input
}
