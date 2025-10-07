package main

import "fmt"

func main() {
	var n, odd int
	fmt.Println("Enter the number")
	fmt.Scanln(&n)
	for i := 1; i < n; i++ {
		if i%2 != 0 {
			odd += i
		}

	}
	fmt.Println("sum of odd :", odd)
}
