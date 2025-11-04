package main

import "fmt"

// Step 1: Struct
type Calculator struct{}

// Step 2: Method
func (c Calculator) Add(a, b int) int {
	return a + b
}

// Step 3: Interface
type Adder interface {
	Add(a, b int) int
}

// Step 4: Function expecting interface
func CalculateSum(a, b int, adder Adder) {
	result := adder.Add(a, b)
	fmt.Println("Sum:", result)
}

// Step 5: Main
func main() {
	CalculateSum(10, 20, Calculator{}) // passing struct directly
}
