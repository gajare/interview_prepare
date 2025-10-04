package main

import "fmt"

func main() {
	var principal, interast, year int
	fmt.Println("Enter principal amount")
	fmt.Scanln(&principal)
	fmt.Println("Enter rate of interest")
	fmt.Scanln(&interast)
	fmt.Println("Enter the years")
	fmt.Scanln(&year)
	simple_interest := (principal * interast * year) / 100
	fmt.Println("Simple interest is :", simple_interest)
}
