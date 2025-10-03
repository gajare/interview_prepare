package main

import "fmt"

func main() {
	var n, count int
	fmt.Println("Enter the number")
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Print("\t", i)
			count++
		}
	}
	if count < 3 {
		fmt.Println("\nThis is prime number :", n)
	}
}
