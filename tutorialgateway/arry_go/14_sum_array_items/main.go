package main

import "fmt"

func main() {
	s1 := []int{1, 6, 9, 2, 4, 6, 7, 8}
	sum := 0
	for i := 0; i < len(s1); i++ {
		sum += s1[i]
	}
	fmt.Println("some of all item in array", sum)
}
