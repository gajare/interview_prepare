package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n string
	fmt.Println("enter number")
	fmt.Scanln(&n)
	num, err := strconv.Atoi(n)
	if err != nil {
		fmt.Println("n is not integer :", err)
		return
	}
	if num < 1 {
		fmt.Println("this is not positive number")
		return
	}
	for i := num; i > 0; i-- {
		fmt.Println(i, "\t")
	}
}
