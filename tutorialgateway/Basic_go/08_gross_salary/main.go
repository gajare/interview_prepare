package main

import "fmt"

func main() {
	var CTC, HRA, DA, grossSalry float64
	fmt.Println("Enter CTC")
	fmt.Scanln(&CTC)
	HRA, DA = GetHraDa(CTC)
	fmt.Println("hra :", HRA)
	fmt.Println("da :", DA)
	grossSalry = CTC + HRA + DA
	fmt.Printf("grocess salary : %f\n", grossSalry)
}

func GetHraDa(CTC float64) (float64, float64) {
	var HRA, DA float64
	if CTC < 10000 {
		HRA = CTC * 0.08
		DA = CTC * 0.1
	} else if CTC < 20000 {
		HRA = CTC * 0.16
		DA = CTC * 0.2
	} else {
		HRA = CTC * 0.24
		DA = CTC * 0.3
	}
	return HRA, DA
}
