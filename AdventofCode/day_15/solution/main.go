package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := readInput()
	inputArr := strings.Split(input, ",")
	hashSum := 0
	for _, v := range inputArr {
		sum := hashUnEditedInput(v)
		hashSum += sum
	}
	fmt.Println(hashSum)
	fmt.Println("=---------part2---------=")
	part2()
}

func hashUnEditedInput(input string) int {
	currentHash := 0
	for i := 0; i < len(input); i++ {
		currentHash = hashChar(string(input[i]), currentHash)
	}
	return currentHash
}

func hashingInput(input string) (string, int) {
	currentHash := 0
	isDash := strings.Contains(input, "-")
	if isDash {
		input = strings.Split(input, "-")[0]
	} else {
		input = strings.Split(input, "=")[0]
	}
	for i := 0; i < len(input); i++ {
		currentHash = hashChar(string(input[i]), currentHash)
	}
	return input, currentHash
}

func hashChar(char string, currentHash int) int {
	currentHash = currentHash + int(rune(char[0]))
	currentHash = currentHash * 17
	return currentHash % 256
}

func readInput() string {
	inputByte, err := os.ReadFile("../input/prod")
	if err != nil {
		panic(err)
	}
	inputByte = inputByte[:len(inputByte)-1]
	inputStr := string(inputByte)
	return inputStr
}
