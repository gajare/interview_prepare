package main

import "fmt"

func main() {
	var i int
	s1 := []int{1, 6, 9, 2, 4, 6, 7, 8}
	key := 8

	flag := false
	for i = 0; i < len(s1); i++ {
		if key == s1[i] {
			flag = true
			break
		}
	}
	if flag {
		fmt.Println("found item:", s1[i], "at index :", i)
	}
}
