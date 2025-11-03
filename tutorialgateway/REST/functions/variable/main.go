package main

import "fmt"

func main() {
	Add := func(a, b int) int {
		return a + b
	}
	var Sub func(int, int) int
	Sub = func(a, b int) int {
		return a - b
	}
	fmt.Println("addition :", Add(10, 10))
	fmt.Println("subtraction :", Sub(10, 5))
}
