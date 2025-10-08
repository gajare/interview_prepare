package main

import "fmt"

func main() {

	var rows, columns int
	fmt.Println("Enter nonber of row")
	fmt.Scanln(&rows)
	fmt.Println("Enter number of coloums")
	fmt.Scanln(&columns)
	matrix := make([][]int, rows)
	fmt.Println("Enter matrix interchange diagonal")
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, columns)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Scanln(&matrix[i][j])
		}
	}
	fmt.Println("matrix :", matrix)
	if rows == columns {
		for i := 0; i < rows; i++ {
			temp := matrix[i][i]
			matrix[i][i] = matrix[i][rows-i-1]
			matrix[i][rows-i-1] = temp
		}
	}
	fmt.Println("matrix interchange")
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Print("\t", matrix[i][j])
		}
		fmt.Println("")
	}
}
