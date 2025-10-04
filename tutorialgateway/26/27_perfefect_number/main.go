package main

import "fmt"

func main() {
	var divisible, n int
	fmt.Println("Enter n number")
	fmt.Scanln(&n)
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			divisible += i
		}
	}
	if n == divisible {
		fmt.Println("perfect number")
	} else {
		fmt.Println("not perfect number")
	}
}
