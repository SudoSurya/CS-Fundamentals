package main

import (
	"fmt"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	var arr [3]int
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])
	fmt.Println(len(arr))

	for i, v := range arr {
		fmt.Printf("%d %d\n", i, v)
	}
	println("--------------------------------")
	var newArr [3]int = [3]int{1, 2, 4}
	for i, v := range newArr {
		fmt.Printf("%d %d\n", i, v)
	}
	println("--------------------------------")

	symbols := [...]string{
		USD: "$",
		EUR: "€",
		GBP: "£",
		RMB: "¥",
	}
	fmt.Println(USD, symbols[USD])

	for i, v := range symbols {
		fmt.Printf("%d %s\n", i, v)
	}
	println("--------------------------------")
	r := [...]int{99: -1}
	fmt.Println(r[99])
	for i, v := range r {
		fmt.Printf("%d %d\n", i, v)
	}
	fmt.Printf("%x", newArr)
	fmt.Println("------------------------------")
	exSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	exSlice1 := []int{1, 2, 3, 5, 5, 6, 7, 8, 9}
	fmt.Println(equal(exSlice, exSlice1))
	nilSlice := []int{}
	if nilSlice == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
    // NOTE: slice is a reference type
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"]++
	ages["charlie"]++
	ages["charlie"]++
	ages["charlie"]++
	ages["charlie"]++
	ages["charlie"]++
	ages["charlie"]++
    fmt.Println(ages["charlie"])
	fmt.Println("------------------------------")
	newAges := map[string]int{
		"alice":   31,
		"charlie": 34,
		"bob":     20,
	}
    for name, age := range newAges {
        fmt.Printf("%s\t%d\n", name, age)
    }
	fmt.Println("------------------------------")
    delete(newAges, "bob")
}

func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
