package main

import "fmt"

func main() {
	var n, sum int
	fmt.Println("Enter number")
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Println("Sum of natural number is:", sum)
	fmt.Println("Average of naturak number is:", sum/n)
}
