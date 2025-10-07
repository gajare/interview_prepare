package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the number")

	for fmt.Scanln(&n); n > 10; n /= 10 {
	}

	fmt.Println("n:", n)
}
