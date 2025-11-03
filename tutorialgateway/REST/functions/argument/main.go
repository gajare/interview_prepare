package main

import "fmt"

func Calculate(a int, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func main() {
	Add := func(a int, b int) int {
		return a + b
	}
	Sub := func(a, b int) int {
		return a - b
	}
	fmt.Println("Addition :", Calculate(10, 10, Add))
	fmt.Println("Addition :", Calculate(10, 5, Sub))
}
