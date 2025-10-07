package main

import "fmt"

func main() {
	s1 := []int{-1, -2, -4, 5, 1, 4, 6, 8, 9, 10}
	var posCount, negCount int
	for _, v := range s1 {
		if v >= 0 {
			posCount++
		} else {
			negCount++
		}
	}
	fmt.Println("positive count :", posCount)
	fmt.Println("negetive count :", negCount)
}
