package main

import "fmt"

func main() {
	var (
		num1     int
		num2     int
		addition int
	)
	fmt.Println("Please enter num1:")
	fmt.Scanln(&num1)
	fmt.Println("Please enter num2:")
	fmt.Scanln(&num2)
	addition = Add(num1, num2)
	fmt.Println("Addition is :", addition)
}

func Add(num1, num2 int) int {
	return num1 + num2
}
