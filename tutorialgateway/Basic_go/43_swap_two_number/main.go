package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Println("Enter two number ")
	fmt.Scanln(&num1, &num2)
	fmt.Println("num1", num1, "\tnum2", num2)
	num1, num2 = num2, num1
	fmt.Println("num1", num1, "\tnum2", num2)
}
