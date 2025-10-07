package main

import (
	"fmt"
	"math"
)

func main() {
	var n, p float64
	fmt.Println("Enter number")
	fmt.Scanln(&n)
	fmt.Println("Enter power number")
	fmt.Scanln(&p)
	res := math.Pow(n, p)
	fmt.Println("result :", res)
}
