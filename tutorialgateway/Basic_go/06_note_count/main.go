package main

import "fmt"

var (
	notes  []int
	amount int
)

func main() {
	notes = []int{500, 100, 50, 20, 10, 5, 2, 1}
	total := 0
	fmt.Println("Enter amount")
	fmt.Scanln(&amount)
	arr := GetNotes(notes, amount)
	for i := 0; i < len(notes); i++ {

		if arr[i] > 0 {
			fmt.Println(notes[i], " : ", arr[i])
			total += notes[i] * arr[i]
		}
	}
	fmt.Println("total amount :", total)
}

func GetNotes(notes []int, amount int) [10]int {
	var arr [10]int

	for i := 0; i < len(notes); i++ {
		arr[i] = amount / notes[i]
		amount = amount % notes[i]
	}

	return arr

}
