package main

import "fmt"

func main() {
	var n int

	fmt.Println("Enter this size of matrix")
	n = 2
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n) // allocate each row
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scanln(&matrix[i][j])
		}
	}
	determinat := (matrix[0][0] * matrix[1][1]) - (matrix[0][1] * matrix[1][0])
	fmt.Println("Determinat of 2x2 matrix :", determinat)
}
