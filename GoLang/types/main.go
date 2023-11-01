package main

import (
	"fmt"
	"log"
	"os"
)

func a() {}

type Celsius float64
type Fahrenheit float64

var cwd string

func init() {
	var err error
	cwd, err = os.Getwd() // NOTE: wrong!
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	log.Printf("Working directory = %s", cwd) //ERROR
}
func main() {
	var a Celsius = 1.0
	var b Fahrenheit = 1.0
	println(a == Celsius(b))
	x := "hello"
	for _, x := range x {
		x := x + 'a' - 'A'
		fmt.Printf("%c", x)
	}
	var apples int32 = 1
	var oranges int16 = 2
    con := int(apples) + int(oranges)
    fmt.Println(con)
}
