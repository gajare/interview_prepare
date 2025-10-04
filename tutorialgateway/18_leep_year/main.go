package main

import "fmt"

func main() {
	var year int
	fmt.Println("Enter the year")
	fmt.Scanln(&year)
	if (year%400 == 0) || (year%4 == 0 && year != 100) {
		fmt.Println("This is leep year")
	} else {
		fmt.Println("This is not leep year")
	}
}
