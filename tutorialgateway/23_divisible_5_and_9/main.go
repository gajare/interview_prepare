package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter a number")
	fmt.Scanln(&n)
	if n%5 == 0 && n%11 == 0 {
		fmt.Println("number is divisible by 5 and 9")
	} else {
		fmt.Println("number is not divisible by 5 and 9")
	}
}
