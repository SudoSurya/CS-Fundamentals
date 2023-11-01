package main

import (
	"fmt"
	"os"
	"path"
	"unicode"
)

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

	f, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println(err)
	}
	// buf := make([]byte,1024)

	// type assertions
	var nums *int = new(int)
	*nums = 222
	nums2 := *nums
	fmt.Println("num2", nums2)
	content := make([]byte, 1024)
	contentLen, err := f.Read(content)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(contentLen)
	fmt.Println(string(content))
	var smalln uint8 = 1
	var bign uint16 = 1
	fmt.Printf("%08b", smalln)
	fmt.Println(smalln == uint8(bign))
	fmt.Println("Hello, playground\a\n  new line content")
	raw := `Hello, playground\n new line content`
	fmt.Println(raw)
	var bl = unicode.IsDigit('1')
	println(bl)
	fmt.Println("rune", string(0x4eac))
	fmt.Println("rune", string(0x4eac))
	fmt.Println(basename("a/b/c.go"))
	mi := basename("/home/kurama/Documents/Personal-Projects/CS-Fundamentals/GoLang/variables/main.go")
	fmt.Println(mi)
}

func g() {
	y := new(int)
	*y = 1
	fmt.Println(y)
	// y doesnt escape from g so it is allocated on the stack
}

func Gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
func basename(s string) string {
	// Discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// Preserve everything before last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
