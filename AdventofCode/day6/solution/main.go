package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var numberRe = regexp.MustCompile(`\d+`)

func main() {
	contents := readInput()
	var raceTimings []string
	var raceDistance []string
	for _, line := range contents {
		raceDetails := numberRe.FindAllString(strings.Split(line, ":")[1], -1)
		switch strings.Split(line, ":")[0] {
		case "Time":
			raceTimings = append(raceTimings, strings.Join(raceDetails, ""))
		case "Distance":
			raceDistance = append(raceDistance, strings.Join(raceDetails, ""))
		}
	}
    fmt.Println(raceTimings)
    fmt.Println(raceDistance)
    wins := []int{}
	for i := 0; i < len(raceTimings); i++ {
		tempRace := ints(raceTimings[i])
        myRace := 0
        maxCount := []int{}
		for tempRace > 0 {
            currRace := tempRace * myRace
            if currRace > ints(raceDistance[i]) {
                maxCount = append(maxCount, myRace)
            }
            myRace++
            tempRace--
		}
        fmt.Println(maxCount)
        wins = append(wins, len(maxCount))
	}
    fmt.Println("wins",mulArr(wins))
}
func ints(input string) int {
	val, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return val
}

func mulArr(input []int) int {
    mul := 1
    for _, val := range input {
        mul *= val
    }
    return mul
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
