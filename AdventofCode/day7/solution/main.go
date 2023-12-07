package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var cardRanks = map[string]int{
	"J": 0,
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"Q": 11,
	"K": 12,
	"A": 13,
}

type Card struct {
	cardType string
	bidValue int
}

func main() {
	input := readInput()
	var CardsMaps = make(map[int][]Card)
	for _, line := range input {
		card, bid := strings.Split(line, " ")[0], strings.Split(line, " ")[1]
		tempCard := Card{cardType: card, bidValue: stringToInt(bid)}
        fmt.Println("tempCard", tempCard)
        fmt.Println("findCardType(card)", findCardType(card))
		CardsMaps[findCardType(card)] = append(CardsMaps[findCardType(card)], tempCard)
	}
	var sortedCards []Card
	for i := 1; i <= 7; i++ {
		if len(CardsMaps[i]) == 0 {
			continue
		}
		fmt.Println("CardsMaps[i]", CardsMaps[i])
		getSoredCards(&sortedCards, CardsMaps[i])
	}
	totalWinnings := 0
	for i := 0; i < len(sortedCards); i++ {
		fmt.Println(i+1, " { ", sortedCards[i].cardType, " , ", sortedCards[i].bidValue, " }")
		totalWinnings += sortedCards[i].bidValue * (i + 1)
	}
	fmt.Println(totalWinnings)
}

func getSoredCards(sortedCards *[]Card, cards []Card) {
	fmt.Println("unsorted cards", cards)
	sortCards(&cards)
	fmt.Println("sorted cards", cards)
	for i := 0; i < len(cards); i++ {
		*sortedCards = append(*sortedCards, cards[i])
	}
}

func compareCards(card1 Card, card2 Card) bool {
	for i := 0; i < len(card1.cardType); i++ {
		if cardRanks[string(card1.cardType[i])] > cardRanks[string(card2.cardType[i])] {
			return false
		}
		if cardRanks[string(card1.cardType[i])] < cardRanks[string(card2.cardType[i])] {
			return true
		}
	}
	return false
}

func findCardType(card string) int {
	cardNumber := getCard(card)
	switch cardNumber {
	case "5":
		return 7
	case "14":
		return 6
	case "23":
		return 5
	case "113":
		return 4
	case "122":
		return 3
	case "1112":
		return 2
	case "11111":
		return 1
	}
	return 0
}

func getCard(card string) string {
	// if all elements are same
	var cardMap = make(map[string]int)
	jCount := 0
	for i := 0; i < len(card); i++ {
		if card[i] == 'J' {
			jCount++
			continue
		}
		cardMap[string(card[i])]++
	}
    cardMap[findMaxKey(cardMap)] += jCount
	cardType := ""
	for _, v := range cardMap {
		cardType += intToString(v)
	}
	return sortString(cardType)
}

func intToString(inputNum int) string {
	str := strconv.Itoa(inputNum)
	return str
}
func stringToInt(inputStr string) int {
	num, err := strconv.Atoi(inputStr)
	if err != nil {
		fmt.Println(err)
	}
	return num
}

func sortString(s string) string {
	// Convert the string to a slice of runes
	runes := []rune(s)

	// Use the sort package to sort the slice
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	// Convert the sorted slice back to a string
	return string(runes)
}
func findMaxKey(hashmap map[string]int) string {
	var maxKey string
	maxValue := 0

	// Iterate over the hashmap to find the key with the maximum value
	for key, value := range hashmap {
		if value > maxValue {
			maxKey = key
			maxValue = value
		}
	}

	return maxKey
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
