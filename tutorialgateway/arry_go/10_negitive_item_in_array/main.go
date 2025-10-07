package main

import "fmt"

func main() {
	s1 := []int{-1, 3, -3, 4, 2, 0, -5}
	fmt.Println("Negitive number in array")
	for _, v := range s1 {
		if v < 0 {
			fmt.Print("\t", v)
		}
	}
}
