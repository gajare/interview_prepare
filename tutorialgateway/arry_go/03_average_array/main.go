package main

import "fmt"

func main() {
	s1 := []int{20, 5, 1, 8, 9, 5, 20, 30}
	var sum int
	for _, v := range s1 {
		sum += v
	}
	fmt.Println("sum :", sum)
	fmt.Println("avg of array s1 :", (sum / len(s1)))
}
