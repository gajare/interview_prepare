package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Println("Enter two numbers")
	fmt.Scanln(&num1, &num2)
	if num1 > num2 {
		fmt.Println("num1 is greater", num1)
	} else {
		fmt.Println("num2 is greater", num2)
	}
}
