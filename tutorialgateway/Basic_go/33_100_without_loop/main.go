package main

import "fmt"

func main() {
	var c = 1
	printNum(c)
}
func printNum(n int) {
	if n <= 100 {
		fmt.Println(n)
		printNum(n + 1)
	}
}
