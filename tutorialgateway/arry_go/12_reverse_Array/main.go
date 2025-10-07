package main

import (
	"fmt"
	"sort"
)

func main() {
	s1 := []int{1, 4, 9, 2, 4, 2, 6, 8}
	fmt.Println("s1 :", s1)
	fmt.Println("reverse array")
	for i := len(s1) - 1; i >= 0; i-- {
		fmt.Print("\t", s1[i])
	}
	fmt.Println("\n Descending order")
	sort.Ints(s1)
	for i := len(s1) - 1; i >= 0; i-- {
		fmt.Print("\t", s1[i])
	}

}
