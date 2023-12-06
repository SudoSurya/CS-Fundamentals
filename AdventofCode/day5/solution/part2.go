package main

import (
	"fmt"
	"strings"
)

func (s *Seeds) pop() []int {
	if len(*s) == 0 {
		return nil
	}

	lastIndex := len(*s) - 1
	poppedElement := (*s)[lastIndex]
	*s = (*s)[:lastIndex]

	return poppedElement
}

type Seeds [][]int

func part2() {
	contents := readInput()
	var seeds Seeds
	for _, line := range contents {

		items := strings.Split(line, "\n")
		switch strings.Split(items[0], ":")[0] {
		case "seeds":
			seedItemsArr := numberRe.FindAllString(items[0], -1)

			for i := 0; i < len(seedItemsArr); i = i + 2 {
				seeds = append(seeds, []int{ints(seedItemsArr[i]), ints(seedItemsArr[i+1]) + ints(seedItemsArr[i])})
			}
		}
	}
	for _, block := range contents[1:] {
		tempRanges := [][]int{}
		for _, line := range strings.Split(block, "\n")[1:] {
			items := numberRe.FindAllString(line, -1)
			tempRanges = append(tempRanges, []int{ints(items[0]), ints(items[1]), ints(items[2])})
		}
		fmt.Println(tempRanges)
		newres := []int{}

		for len(seeds) > 0 {
			popSeed := seeds.pop()
			start := popSeed[0]
			end := popSeed[1]
			for _, tempRange := range tempRanges {
				os := max(start, tempRange[1])
				oe := min(end, tempRange[0]+tempRange[2])
				if os < oe {
					newres = append(newres, []int{os - tempRange[1] + tempRange[0], oe - tempRange[1] + tempRange[0]}...)
					if os > start {
						seeds = append(seeds, []int{start, os})
					}
					if end > oe {
						seeds = append(seeds, []int{oe, end})
					}
				} else {
					break
				}
			}
            fmt.Println("seeds", seeds)
            newres = append(newres, []int{start, end}...)
		}
		fmt.Println("newres", newres)
		seeds = append(seeds, newres)
	}
	fmt.Println("seeds", seeds)
}
