package main

import "fmt"

func main() {
	var units, amount, sCharage float64
	units = 220
	if units <= 50 {
		amount = units * 2.6
		sCharage = 25
	} else if units <= 100 {
		amount = 130 + ((units - 50) * 3.25)

		sCharage = 35
	} else if units <= 200 {
		amount = ((130 + 162.5) + ((units - 100) * 5.65))

		sCharage = 45
	} else {
		amount = ((130 + 162.5 + 526) + ((units - 200) * 7.75))
		sCharage = 55
	}
	fmt.Println("amount :", amount+sCharage)
	fmt.Println("s charge :", sCharage)
}
