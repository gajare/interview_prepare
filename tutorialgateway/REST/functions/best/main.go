package main

import "fmt"

func Cal(op string) func(int) func() {
	return func(i int) func() {
		if op == "square" {
			return func() {
				fmt.Println("square", i*i)
			}
		} else if op == "cube" {
			return func() {
				fmt.Println("cube :", i*i*i)
			}
		} else {
			return func() {
				fmt.Println("Invalide operation")
			}
		}
	}
}

func main() {
	operation := Cal("square")
	res := operation(5)
	res()
	Cal("cube")(2)()
}

//func RequireRole(requiredRole string) func(next http.Handler) http.Handler
