package main

import (
	"fmt"
)

// Function to calculate the Greatest Common Divisor (GCD)
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Function to calculate the Least Common Multiple (LCM)
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// Function to calculate the LCM of multiple numbers
func calculateLCM(numbers []int) int {
	result := 1
	for _, num := range numbers {
		result = lcm(result, num)
	}
	return result
}

func LCM(numbers []int) {
	// Example: Calculate LCM of 4, 6, and 8
	result := calculateLCM(numbers)

	fmt.Printf("LCM of %v is %d\n", numbers, result)
}
