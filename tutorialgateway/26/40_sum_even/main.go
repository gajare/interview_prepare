package main

import "fmt"

func main() {
	var n, sum int
	fmt.Println("Enter the number")
	fmt.Scanln(&n)
	for i := 0; i <= n; i++ {
		if i%2 == 0 {
			sum = sum + i
		}
	}
	fmt.Println("Sum of even numbers", sum)
}
