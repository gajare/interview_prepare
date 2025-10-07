package main

import "fmt"

func main() {

	var rem, rev, n, temp int
	fmt.Println("Enter a number")
	fmt.Scanln(&n)
	temp = n
	for n > 0 {
		rem = n % 10
		rev = (rev * 10) + rem
		n = n / 10
	}
	if temp == rev {
		fmt.Println("number is palindrom")
	} else {
		fmt.Println("number is not")
	}
}
