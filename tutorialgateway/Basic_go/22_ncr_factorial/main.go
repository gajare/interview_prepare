package main

import "fmt"

func main() {
	var n, c int
	fmt.Println("Enter  total student")
	// fmt.Scanln(&n)
	n = 5
	fmt.Println("Enter choose student")
	// fmt.Scanln(&c)
	c = 2
	ncr := factorial(n) / (factorial(c) * factorial(n-c))
	fmt.Println("", c)
	fmt.Println("NCR :", ncr)
}

func factorial(n int) int {
	var fact = 1
	for i := 1; i <= n; i++ {
		fact = fact * i
	}
	return fact
}
