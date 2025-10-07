package main

import "fmt"

func main() {
	var n, fact int
	fmt.Println("Enter the N number to get factorial")
	fmt.Scanln(&n)
	fact = Factorial(n)
	fmt.Printf("factorial of %d is %d", n, fact)
}

func Factorial(n int) int {
	if n == 1 {
		return 1
	}
	return Factorial(n-1) * n

}
