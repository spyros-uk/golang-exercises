package main

import (
	"fmt"
)
func main() {
	sum := Add(3, 4)
	fmt.Printf("Sum of 3 + 4 is %v (%T)", sum, sum)
}

func Add(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}

func Sub(a, b int) int  {
	return a - b
}