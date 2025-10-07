package main

import "fmt"

func main() {
	val, count := 4311232, 0
	for val > 0 {
		count++
		val /= 10
	}
	fmt.Println("total digits are:", count)
}
