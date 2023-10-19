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

	fmt.Println(&a == &b)
}
func g() {
	y := new(int)
	*y = 1
    // y doesnt escape from g so it is allocated on the stack
}
