package main

import "fmt"

func main() {
	s1 := []int{1, 3, 6, 2, 4, 8, 2, 9, 10}
	odd := []int{}
	even := []int{}
	for _, v := range s1 {
		if v%2 == 0 {
			even = append(even, v)
		}
		if v%2 != 0 {
			odd = append(odd, v)
		}
	}
	fmt.Println("odd aray :", odd)
	fmt.Println("even array :", even)
}
