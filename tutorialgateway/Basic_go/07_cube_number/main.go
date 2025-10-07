package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter a number")
	fmt.Scanln(&num)
	cube := num * num * num
	fmt.Println("cube of a number :", cube)
}
