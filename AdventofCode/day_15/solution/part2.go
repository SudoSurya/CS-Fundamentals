package main

import (
	"fmt"
	"strconv"
	"strings"
)

type focusMap struct {
	hashcode string
	focalLen int
}

func part2() {
	input := readInput()
	inputArr := strings.Split(input, ",")

	sequence := map[int][]focusMap{}

	for _, v := range inputArr {
		code, hashcode := hashingInput(v)
		if strings.Contains(v, "=") {
			fl, _ := strconv.Atoi(strings.Split(v, "=")[1])
            focalLen := fl

			isCodeExist := false
			for _, v := range sequence[hashcode] {
				if v.hashcode == code {
					isCodeExist = true
				}
			}
			if !isCodeExist {
				sequence[hashcode] = append(sequence[hashcode], focusMap{code, focalLen})
			} else {
				for i, v := range sequence[hashcode] {
					if v.hashcode == code {
						sequence[hashcode][i].focalLen = focalLen
					}
				}
			}
		}
		if strings.Contains(v, "-") {
			isCodeExist := false
			for _, v := range sequence[hashcode] {
				if v.hashcode == code {
					isCodeExist = true
				}
			}
			if isCodeExist {
				indexTodelete := -1
				for i, v := range sequence[hashcode] {
					if v.hashcode == code {
						indexTodelete = i
					}
				}
				if indexTodelete != -1 {
					sequence[hashcode] = append(sequence[hashcode][:indexTodelete], sequence[hashcode][indexTodelete+1:]...)
				}
			}
		}
	}
    focusPoint := 0
	for k, v := range sequence {
		for k1, v2 := range v {
            focusPoint += ((k+1) * (k1+1) * v2.focalLen)
		}
	}
    fmt.Println(focusPoint)
}
