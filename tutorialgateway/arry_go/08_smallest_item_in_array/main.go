package main

import "fmt"

type Small struct {
	value int
	index int
}
type Largest struct {
	value int
	index int
}

func main() {
	s1 := []int{10, 3, 6, 9, 1, 2, 50, 7}
	smallest := Small{value: 0, index: 0}
	largest := Largest{value: 0, index: 0}
	for i := 0; i < len(s1); i++ {
		if smallest.value > s1[i] {
			smallest.value = s1[i]
			smallest.index = i
		}
		if largest.value <= s1[i] {
			largest.value = s1[i]
			largest.index = i
		}
	}
	fmt.Println("smallest item in array :", smallest.value, " at index :", smallest.index)
	fmt.Println("Largest item in array :", largest.value, " at index :", largest.index)
}
