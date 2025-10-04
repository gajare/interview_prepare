package main

import "fmt"

func main() {
	var n = 123
	var rem, rev int
	// rev = 1

	fmt.Println("n:", n)
	for n > 0 {
		rem = n % 10
		rev = (rev * 10) + rem
		n = n / 10
	}
	fmt.Println("reverse number :", rev)
}
