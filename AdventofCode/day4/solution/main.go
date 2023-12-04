package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

var (
	numberRe = regexp.MustCompile(`\d+`)
)

type Card struct {
	Copy     int
	Macthes  int
	Original int
}

func main() {
	contents, err := ReadInput()
	if err != nil {
		panic(err)
	}
	points := 0
	ScratchCards := []Card{}
	for _, line := range contents {
		cardsSplit := strings.Split(line, ":")
		cards := strings.Split(cardsSplit[1], "|")
		points += calculateDoubled(
			countCommonElements(
				numberRe.FindAllString(cards[0], -1),
				numberRe.FindAllString(cards[1], -1),
			))

		tempCard := Card{
			Copy:     0,
			Macthes:  countCommonElements(numberRe.FindAllString(cards[0], -1), numberRe.FindAllString(cards[1], -1)),
			Original: 1,
		}
		ScratchCards = append(ScratchCards, tempCard)
	}

	totalScrathCards := 0
	for idx, card := range ScratchCards {
		totalScrathCards += card.Copy + card.Original
		if card.Macthes == 0 {
			continue
		}
		for i := 0; i < card.Macthes; i++ {
			ScratchCards[idx+1+i].Copy += (card.Copy + card.Original)
		}
	}
	fmt.Println("points", points)
	fmt.Println("totalScrathCards", totalScrathCards)
}
func calculateDoubled(input int) int {
	return int(math.Pow(2, float64(input-1)))
}
func countCommonElements(array1, array2 []string) int {
	set := make(map[string]bool)
	var count int

	for _, element := range array1 {
		set[element] = true
	}

	for _, element := range array2 {
		if set[element] {
			count++
		}
	}

	return count
}

func ReadInput() ([]string, error) {
	inputByte, err := os.ReadFile("../input/prod")
	if err != nil {
		return nil, err
	}
	inputByte = inputByte[:len(inputByte)-1]
	inputStr := string(inputByte)
	input := strings.Split(inputStr, "\n")
	return input, nil
}
