package main

import "fmt"

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
}
