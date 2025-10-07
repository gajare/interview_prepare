package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var oddCount, evenCount int
	for i := 0; i < len(s1); i++ {
		if s1[i]%2 == 0 {
			oddCount++
		} else {
			evenCount++
		}
	}
	fmt.Println("odd count :", oddCount)
	fmt.Println("even count :", evenCount)
}
