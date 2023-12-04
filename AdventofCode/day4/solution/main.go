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

	fmt.Println("before", ScratchCards)
    sum := 0
	for idx, card := range ScratchCards {
		if card.Macthes == 0 {
			continue
		}
		fmt.Println("idx", idx, "card", card)
		for i := 0; i < card.Macthes; i++ {
			key := idx + 1 + i
			cardCopy := ScratchCards[key]
			cardCopy.Copy = cardCopy.Copy + (card.Copy + card.Original)
			ScratchCards[key] = cardCopy
		}
		fmt.Println("idx", idx, "card", card)
	}
	fmt.Println("after", ScratchCards)
	for _, card := range ScratchCards {
		sum += card.Copy + card.Original
	}
	fmt.Println("sum", sum)

}
func calculateDoubled(input int) int {
	return int(math.Pow(2, float64(input-1)))
}
func countCommonElements(array1, array2 []string) int {
	set := make(map[string]bool)
	var count int

	// Create a set from array1
	for _, element := range array1 {
		set[element] = true
	}

	// Check for common elements in array2
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
