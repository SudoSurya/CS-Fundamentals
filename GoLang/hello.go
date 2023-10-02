package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")

	/* s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s) */

	input := bufio.NewScanner(os.Stdin)

	input.Scan()
	println(input.Text())
    var output bool
    println(output)
}

