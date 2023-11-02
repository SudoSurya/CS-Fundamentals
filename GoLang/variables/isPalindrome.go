package main

import (
	"fmt"
	"strings"
)

func init() {
	fmt.Println("init")

	println("----------------------")
	println("----------------------")
	println("----------------------")
	res := isPalindromes("A man, a plan, a canal: Panama")
	fmt.Println("res", res)
	println("----------------------")
}

func isPalindromes(s string) bool {
	if len(s) < 2 || s == " " {
		return true
	}
	s = strings.ToLower(s)
	fmt.Println(s)

	var right int = 0
	var left int = len(s) - 1

	for left > right {
		if IsChar(s[left]) && IsChar(s[right]) {
			if s[left] == ' ' {
				left--
				continue
			}
			if s[right] == ' ' {
				right++
				continue
			}
			if s[left] != s[right] {
				fmt.Println("false")
				return false
			}
			left--
			right++
		} else if !IsChar(s[left]) {
			left--
		} else if !IsChar(s[right]) {
			right++
		}
	}
	fmt.Println("true")
	return true
}

func IsChar(c byte) bool {
	return (c >= 'a' && c <= 'z' || c >= '0' && c <= '9')
}
