package main

import "fmt"

func Calculate(s1 string) func(int, int) int {
	if s1 == "add" {
		return func(i1, i2 int) int {
			return i1 + i2
		}
	}
	return nil
}

func main() {
	getAdd := Calculate("add")
	fmt.Println("addition :", getAdd(10, 10))
}
