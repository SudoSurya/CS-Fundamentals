package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	/* s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s) */

	/* input := bufio.NewScanner(os.Stdin)

		input.Scan()
		println(input.Text())
	    var output bool
	    println(output) */

	x := 0
	p := &x
	if p != nil {
		println(*p)
	}
	println(*p)
	println(p)
}
