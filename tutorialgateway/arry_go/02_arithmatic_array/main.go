package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{6, 7, 8, 9, 10}
	fmt.Println("Add\tSub\tMul\tDiv\tMod")
	for i := 0; i < len(s2); i++ {

		fmt.Printf("\n %d\t", s2[i]+s1[i])
		fmt.Printf("%d\t", s2[i]-s1[i])
		fmt.Printf("%d\t", s2[i]*s1[i])
		fmt.Printf("%d\t", s2[i]/s1[i])
		fmt.Printf("%d\t", s2[i]%s1[i])
	}
}
