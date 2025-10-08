package main

import "fmt"

func main() {
	matrix1 := [][]int{{1, 2}, {1, 2}}
	matrix2 := [][]int{{1, 2}, {1, 2}}
	// Addition := make([][]int, 2)
	fmt.Println(" Matrix 1 ")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print("\t\t", matrix1[i][j])
		}
		fmt.Println()
	}
	fmt.Println(" Martrix 2 ")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print("\t\t", matrix2[i][j])
		}
		fmt.Println()
	}
	fmt.Println("Additional matrix ")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print("\t\t", matrix1[i][j]+matrix2[i][j])
		}
		fmt.Println()
	}
}
