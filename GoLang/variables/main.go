package main

import "fmt"

func main() {

	var p = new(int)
	fmt.Println(*p)
	var foo = new(int)
	fmt.Println(*foo)

	a := 1
	b := 1

	fmt.Println(p == foo)
    g()
	fmt.Println(&a == &b)
	result := Gcd(10, 20)
	fmt.Println(result)
}
func g() {
	y := new(int)
	*y = 1
	// y doesnt escape from g so it is allocated on the stack
}

func Gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
