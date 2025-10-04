package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter a number")
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		if i%2 != 0 {
			fmt.Println(i, "\t")
		}
	}
}
