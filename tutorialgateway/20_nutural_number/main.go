package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var n string
	fmt.Println("Enter number")
	fmt.Scanln(&n)
	num, err := strconv.Atoi(n)
	if err != nil {
		fmt.Println("number is not integer")
	} else if num < 1 {
		fmt.Println("Please enter positive value")
	} else {
		for i := 1; i <= num; i++ {
			fmt.Println(i, "\t")
		}
	}
	fmt.Println("err :", reflect.TypeOf(err))
}
