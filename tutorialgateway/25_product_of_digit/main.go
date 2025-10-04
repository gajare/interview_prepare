package main

import "fmt"

func main() {
	var n, rem, prod int
	fmt.Println("Enter a number")
	fmt.Scanln(&n)
	prod = 1
	for ; n >= 1; n /= 10 {
		rem = n % 10
		prod = prod * rem

	}
	fmt.Println("product of digit :", prod)
}
