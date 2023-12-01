package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input := ReadInput()
	var part1 int
	var part2 int
	for _, val := range input {
		getval := GetCalibration(val, 1)
		getval2 := GetCalibration(val, 2)
		part1 += getval
		part2 += getval2
	}
    fmt.Println("Part 1:", part1)
    fmt.Println("Part 2:", part2)
}

func ReadInput() []string {
	inputByte, err := os.ReadFile("../input/01-prod.text")
	if err != nil {
		panic(err)
	}
	inputByte = inputByte[:len(inputByte)-1]
	inputStr := string(inputByte)
	input := strings.Split(inputStr, "\n")
	return input
}

func GetCalibration(input string, isPart int) int {
	if isPart == 2 {
		input = matchAndReplace(input)
		input = matchAndReplace(input)
	}
	return GetDigits(input)
}

func getLastDigit(input string) string {
	var digit string
	for i := len(input) - 1; i >= 0; i-- {
		char := input[i]
		if unicode.IsDigit(rune(char)) {
			digit = string(char) + digit
			break
		}
	}
	return digit
}

func GetDigits(input string) int {
	num, err := strconv.Atoi(getFirstDigit(input) + getLastDigit(input))
	if err != nil {
		// Handle the error (e.g., log it, return a default value, etc.)
		fmt.Println("Error converting string to int:", err)
		return 0
	}
	return num
}

func getFirstDigit(input string) string {
	var digit string
	for _, char := range input {
		if unicode.IsDigit(char) {
			digit += string(char)
			break
		}
	}
	return digit
}

func matchAndReplace(input string) string {
	re := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|1|2|3|4|5|6|7|8|9)`)
	result := re.ReplaceAllStringFunc(input, func(match string) string {
		return wordToNumber(match)
	})
	return result
}

func wordToNumber(word string) string {
	switch word {
	case "one":
		return "1e"
	case "two":
		return "2o"
	case "three":
		return "3e"
	case "four":
		return "4r"
	case "five":
		return "5e"
	case "six":
		return "6x"
	case "seven":
		return "7n"
	case "eight":
		return "8t"
	case "nine":
		return "9e"
	case "1":
		return "1"
	case "2":
		return "2"
	case "3":
		return "3"
	case "4":
		return "4"
	case "5":
		return "5"
	case "6":
		return "6"
	case "7":
		return "7"
	case "8":
		return "8"
	case "9":
		return "9"
	}
	return word // return the original word if not a known number word
}
