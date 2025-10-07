package main

import "fmt"

func main() {
	var n, g int
	fmt.Println("Please Enter a number")
	n = 786
	g = n
	for g > 9 {
		g = Root(g)
	}
	fmt.Println("generic number is :", g)
}
func Root(n int) int {
	var rem int
	for n > 0 {
		rem += n % 10
		n = n / 10
	}
	return rem
}
