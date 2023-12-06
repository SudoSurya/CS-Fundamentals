package main

import (
	"fmt"
	"strings"
)

func part1() {
	contents := readInput()
	var seedItems []int
	for _, line := range contents {
		items := strings.Split(line, "\n")
		switch strings.Split(items[0], ":")[0] {
		case "seeds":
			seedItemsArr := numberRe.FindAllString(items[0], -1)

			for i := 0; i < len(seedItemsArr); i++ {
				seedItems = append(seedItems, ints(seedItemsArr[i]))
			}

		case "seed-to-soil map":
			convertToAlmanc(&seedItems, &items)
		case "soil-to-fertilizer map":
			convertToAlmanc(&seedItems, &items)
		case "fertilizer-to-water map":
			convertToAlmanc(&seedItems, &items)
		case "water-to-light map":
			convertToAlmanc(&seedItems, &items)
		case "light-to-temperature map":
			convertToAlmanc(&seedItems, &items)
		case "temperature-to-humidity map":
			convertToAlmanc(&seedItems, &items)
		case "humidity-to-location map":
			convertToAlmanc(&seedItems, &items)
		}
	}
	fmt.Println("range", findMinElement(seedItems))

}
