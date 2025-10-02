package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		tenure       float64
		interestRate float64
		principal    float64
		totalAmount  float64
	)
	fmt.Println("Enter the amount needed")
	fmt.Scanln(&principal)
	fmt.Println("Enter the tenure")
	fmt.Scanln(&tenure)
	fmt.Println("Enter interest rate")
	fmt.Scanln(&interestRate)
	compoundIntres, totalAmount := GetCompoundInterest(principal, tenure, interestRate)
	fmt.Printf("Compound interest :%.2f\n", compoundIntres)
	fmt.Printf("Total amount after :%.0f years :%.2f\n", tenure, totalAmount)
}

func GetCompoundInterest(principal, tenure, interestRate float64) (float64, float64) {
	futureValue := principal * math.Pow(1+(interestRate/100), tenure)
	compoundIntres := futureValue - principal
	return compoundIntres, futureValue
}
