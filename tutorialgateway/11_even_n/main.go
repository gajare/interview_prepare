package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter number")
	fmt.Scanln(&n)
	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			fmt.Print(i, "\t")
		}

	}
}
