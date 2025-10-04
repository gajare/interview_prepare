package main

import "fmt"

func main() {
	var price, sale, profit, loss float64
	fmt.Println("Enter product price")
	fmt.Scanln(&price)
	fmt.Println("Enter saling price")
	fmt.Scanln(&sale)
	if price > sale {
		loss = price - sale
		fmt.Printf("Loss :%.2f\n", loss)
	} else if price < sale {
		profit = sale - price
		fmt.Printf("profit :%.2f", profit)
	} else {
		fmt.Println("There is no profit and no loss")
	}
}
