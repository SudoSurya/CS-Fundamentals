package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	fmt.Println("Hello, World!")

	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}

	z := new(int)
    fmt.Println("ss",&z)
	pop := new(int)
    fmt.Println("ss",&pop)
	foo := new(int)
    fmt.Println("ss",&foo)
	bar := new(int)
    fmt.Println("ss",&bar)

	// s, sep := "", ""
	// for& _, arg := range os.Args[1:] {
	// 	s += sep + arg
	// 	sep = " "
	// }
	// fmt.Println(s)

	/* input := bufio.NewScanner(os.Stdin)

		input.Scan()
		println(input.Text())
	    var output bool
	    println(output) */

	x := 0
	p := &x
	// if &p != nil {
	// 	println(*p)
	// }
	// println(p)

	v := increment(p)
	println(v)
	println(x)
}

func increment(inc *int) int {
	*inc++
	return *inc
}
