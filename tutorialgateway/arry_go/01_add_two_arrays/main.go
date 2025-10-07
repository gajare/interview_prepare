package main

import "fmt"

func main() {
	oddArray := []int{1, 3, 5, 7, 9}
	evenArray := []int{2, 4, 6, 8, 10}
	resArray := make([]int, 0, 10)

	for i := 0; i < len(oddArray) || i < len(evenArray); i++ {
		if i == len(oddArray) {
			resArray = append(resArray, evenArray[i])
		}
		if i == len(evenArray) {
			resArray = append(resArray, oddArray[i])
		}
		resArray = append(resArray, oddArray[i]+evenArray[i])
	}
	fmt.Println("resArray :", resArray)
}
