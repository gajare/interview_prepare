package main

import "fmt"

func main() {
	var n, sum int
	fmt.Println("Enter a number")
	fmt.Scanln(&n)
	for sum = 0; n > 0; n /= 10 {
		rem := n % 10
		sum += rem
	}
	fmt.Println("Sum of digits in number is", sum)
}
