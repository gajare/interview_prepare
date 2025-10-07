package main

import "fmt"

func main() {
	s1 := []int{1, 3, 2, 5, 3, 7, 8, 9, 10}
	fmt.Println("All even position")
	for i := 0; i < len(s1); i++ {
		if s1[i]%2 == 0 {
			fmt.Print("\t", s1[i])
		}

	}
	fmt.Println("\nAll odd position ")
	for i := 0; i < len(s1); i++ {
		if s1[i]%2 != 0 {
			fmt.Print("\t", s1[i])
		}
	}
}
