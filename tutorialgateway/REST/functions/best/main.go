package main

import "fmt"

func Cal(op string) func(int) func() {
	return func(i int) func() {
		switch op {
		case "square":
			return func() {
				fmt.Println("square :", i*i)
			}
		case "cube":
			return func() {
				fmt.Println("cube :", i*i*i)
			}
		default:
			return func() {
				fmt.Println("wrong input text")
			}
		}
	}
}

func main() {
	operation := Cal("square")
	res := operation(5)
	res()
	Cal("cube")(2)()
	Cal("Amol")(2)()
	Cal("cube")      // not invoke golang igore return value
	Cal("cube")(2)   // not invoke golang igore return value
	Cal("Amol")(2)() // invoked

}

//func RequireRole(requiredRole string) func(next http.Handler) http.Handler
