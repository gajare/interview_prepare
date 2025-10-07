package main

import (
	"fmt"
	"sort"
)

func main() {
	s1 := []int{1, 3, 6, 8, 11, 45, 2, 6, 8, 9}
	sort.Ints(s1)
	fmt.Println("largest number in array :", s1[len(s1)-1])
	fmt.Println("smallest item in array :", s1[0])

}
