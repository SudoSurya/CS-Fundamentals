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

			fmt.Println("seedItems", seedItems)
		case "seed-to-soil map":
			convertToAlmanc(&seedItems, &items)
			fmt.Println("seedItems", seedItems)
		case "soil-to-fertilizer map":
			convertToAlmanc(&seedItems, &items)
			fmt.Println("seedItems", seedItems)
		case "fertilizer-to-water map":
			convertToAlmanc(&seedItems, &items)
			fmt.Println("seedItems", seedItems)
		case "water-to-light map":
			convertToAlmanc(&seedItems, &items)
			fmt.Println("seedItems", seedItems)
		case "light-to-temperature map":
			convertToAlmanc(&seedItems, &items)
			fmt.Println("seedItems", seedItems)
		case "temperature-to-humidity map":
			convertToAlmanc(&seedItems, &items)
			fmt.Println("seedItems", seedItems)
		case "humidity-to-location map":
			convertToAlmanc(&seedItems, &items)
			fmt.Println("seedItems", seedItems)
		}
	}
	fmt.Println("lowestLocation", findMinElement(seedItems))
    fmt.Println("part1 end")
}
