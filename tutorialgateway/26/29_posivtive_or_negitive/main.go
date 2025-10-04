package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter a number")
	fmt.Scanln(&n)
	if n >= 0 {
		fmt.Println("Number is positive")
	} else {
		fmt.Println("Number is negitive")
	}
}
