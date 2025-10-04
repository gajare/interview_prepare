package main

import "fmt"

func main() {
	var n, odd, even int
	fmt.Println("Enter a number")
	fmt.Scanln(&n)
	for i := 1; i < n; i++ {
		if i%2 == 0 {
			even = even + i
		} else {
			odd = odd + i
		}
	}
	fmt.Println("Sum of odd :", odd)
	fmt.Println("Sum of even :", even)
}
