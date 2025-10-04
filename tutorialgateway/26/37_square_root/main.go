package main

import (
	"fmt"
	"math"
)

func main() {
	var n, i float64
	var flag = false
	fmt.Println("Enter squre number")
	fmt.Scanln(&n)
	for i = 1; i <= n/2; i++ {
		if i*i == n {
			flag = true
			break
		}
	}
	if flag {
		fmt.Println("square root is :", i)
	} else {
		fmt.Println("Not able to get its sqare root")
	}

	fmt.Println("sqare root with using function", math.Sqrt(n))

}
