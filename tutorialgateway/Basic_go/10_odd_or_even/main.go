package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter number")
	fmt.Scanln(&num)
	if num%2 == 0 {
		fmt.Printf("%d num is even", num)
	} else {
		fmt.Printf("%d num is odd", num)
	}
}
