package main

import "fmt"

func main() {
	var n, i int
	fmt.Println("Enter a num")
	fmt.Scanln(&n)
	for i = 1; i <= 10; i++ {
		fmt.Println(n * i)
	}
}
