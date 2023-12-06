package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part2() {
	contents := ReadInput()
	mapRanges := get_input_map(contents)
	seedsCopy := numberRe.FindAllString(strings.Split(contents[0], ":")[1], -1)
	seeds := []int{}
	for i := 0; i < len(seedsCopy); i = i + 2 {
		seeds = append(seeds, ints(seedsCopy[i]), ints(seedsCopy[i+1]))
	}
	fmt.Println(seeds)
	output := seedsToRange(seeds, mapRanges)
	fmt.Println("output", output)
	fmt.Println("end")
}
func seedsToRange(seeds []int, mapRanges map[int][][]int) int {

	for i := findMin(seeds); ; i+=2 {
		location := i

		for j := len(mapRanges); j > 0; j-- {
			location = compute_seed(location, mapRanges[j])
		}

		for k := 0; k < len(seeds); k += 2 {
			x := seeds[k]
			y := seeds[k+1]

			if x <= location && location < x+y {
				return i

			}
		}
	}
}

func compute_seed(current_seed int, mapping [][]int) int {
	for _, m := range mapping {
		destination := m[0]
		start := m[1]
		length := m[2]
		if destination <= current_seed && current_seed < (destination+length) {
			new_seed := start + (current_seed - destination)
			return new_seed
		}

	}
	return current_seed
}

func ReadInput() []string {
	inputByte, err := os.ReadFile("../input/prod")
	if err != nil {
		panic(err)
	}
	inputByte = inputByte[:len(inputByte)-1]
	inputStr := string(inputByte)
	input := strings.Split(inputStr, "\n\n")
	return input
}

func get_input_map(lines []string) map[int][][]int {
	all_maps := make(map[int][][]int)

	for i := 1; i < len(lines); i++ {

		single_map := [][]int{}
		for _, line := range strings.Split(lines[i], "\n") {
			matches := regexp.MustCompile(`-?\d+`).FindAllString(line, -1)
			map_values := []int{}

			for _, match := range matches {
				num, err := strconv.Atoi(match)
				if err == nil {
					map_values = append(map_values, num)
				}
			}
			if len(map_values) > 0 {
				single_map = append(single_map, map_values)
			}
		}

		all_maps[i] = single_map

	}

	return all_maps
}
