package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	result := addSum(5, 3)
	fmt.Println("The sum is:", result)
}

func addSum(a int, b int) int {
	return a + b
}
