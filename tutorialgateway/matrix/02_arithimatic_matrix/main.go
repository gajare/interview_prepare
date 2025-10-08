package main

import "fmt"

func main() {
	matrix1 := [][]int{
		{10, 20},
		{5, 10}}
	matrix2 := [][]int{
		{2, 4},
		{2, 6}}

	fmt.Println("Addition of matrix")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print("\t", matrix1[i][j]+matrix2[i][j])
		}
		fmt.Println()
	}
	fmt.Println("Subtraction of matrix")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print("\t", matrix1[i][j]-matrix2[i][j])
		}
		fmt.Println()
	}
	fmt.Println("Multiplication of matrix")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print("\t", matrix1[i][j]*matrix2[i][j])
		}
		fmt.Println()
	}
	fmt.Println("Division of matrix")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print("\t", matrix1[i][j]/matrix2[i][j])
		}
		fmt.Println()
	}
	fmt.Println("Mod of matrix")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print("\t", matrix1[i][j]%matrix2[i][j])
		}
		fmt.Println()
	}

}
