package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	color  string
	number int
}
type Games []Game

func main() {

	input, err := ReadInput()
	if err != nil {
		panic(err)
	}
	gameData := map[string][]Game{}
	for _, line := range input {
		lineSplit := strings.Split(line, ":")
		gameName := lineSplit[0]
		recordedGames := strings.Split(lineSplit[1], ";")
		gameData[gameName] = []Game{}
		for _, game := range recordedGames {
			for _, indGame := range strings.Split(game, ",") {
				sepGame := strings.Split(indGame, " ")
				number, err := strconv.Atoi(sepGame[1])
				if err != nil {
					panic(err)
				}
				gameData[gameName] = append(
					gameData[gameName],
					Games{Game{sepGame[2], number}}...,
				)

			}
		}
	}
	CubeConudrum(gameData)
	CubeConudrum2(gameData)
}

func CubeConudrum(gameData map[string][]Game) {
	vaildCubes := map[string]int{
		"red":   12,
		"blue":  14,
		"green": 13,
	}
	validGames := []string{}
	for gameName, games := range gameData {
		isGameValid := true
		red, blue, green := 0, 0, 0
		for _, game := range games {
			switch game.color {
			case "red":
				red = game.number
			case "blue":
				blue = game.number
			case "green":
				green = game.number
			}
			if red > vaildCubes["red"] || blue > vaildCubes["blue"] || green > vaildCubes["green"] {
				isGameValid = false
				break
			}
		}
		if isGameValid {
			validGames = append(validGames, gameName)
		}
	}
	GameidxSum := 0
	sumIdx(validGames, &GameidxSum)
    fmt.Println("part1:",GameidxSum)
}

func sumIdx(validGames []string, GameidxSum *int) {
	for _, char := range validGames {
		gameidx := char[5:]
		gameidxInt, err := strconv.Atoi(gameidx)
		if err != nil {
			panic(err)
		}
		*GameidxSum += gameidxInt
	}
}

func CubeConudrum2(gameData map[string][]Game) {
	powerSets := 0
	for _, games := range gameData {
		red, blue, green := 0, 0, 0
		for _, game := range games {
			switch game.color {
			case "red":
				red = max(red, game.number)
			case "blue":
				blue = max(blue, game.number)
			case "green":
				green = max(green, game.number)
			}
		}
		powerCube := red * blue * green
		powerSets += powerCube
	}
    fmt.Println("part2:",powerSets)
}

func ReadInput() ([]string, error) {
	inputByte, err := os.ReadFile("../input/01-prod.txt")
	if err != nil {
		return nil, err
	}
	inputByte = inputByte[:len(inputByte)-1]
	inputStr := string(inputByte)
	input := strings.Split(inputStr, "\n")
	return input, nil
}
