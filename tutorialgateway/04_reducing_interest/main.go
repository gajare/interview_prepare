package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		pricipal   float64
		tenure     float64
		annualRate float64
	)
	fmt.Println("Enter the amount")
	fmt.Scanln(&pricipal)
	fmt.Println("Enter the Tenure")
	fmt.Scanln(&tenure)
	fmt.Println("Enter Intrest percentages")
	fmt.Scanln(&annualRate)
	emi, totalIntrest, totalPayment := GetReducingBalenceEmi(pricipal, tenure, annualRate)
	fmt.Println("EMI :", emi)
	fmt.Println("totalInterest :", totalIntrest)
	fmt.Println("totoal payment :", totalPayment)
}

func GetReducingBalenceEmi(principal, tenure, annualRate float64) (float64, float64, float64) {
	months := tenure * 12
	monthlyRate := annualRate / 12 / 100
	emi := principal * monthlyRate * math.Pow(1+monthlyRate, months) / (math.Pow(1+monthlyRate, months) - 1)
	totalPayment := emi * months
	totalInterest := totalPayment - principal
	return emi, totalPayment, totalInterest
}
