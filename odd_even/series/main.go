package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Odd(ch chan int) {
	defer wg.Done()
	for i := 1; i < 10; i += 2 {
		ch <- i
		fmt.Println(" odd :", i)
		<-ch //wait for receving the value
	}
}

func Even(ch chan int) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		<-ch
		fmt.Println("even :", i)
		ch <- i //send value that waiting for
	}
}
func main() {
	wg.Add(2)
	ch := make(chan int, 5)

	go Even(ch)
	go Odd(ch)
	wg.Wait()
}
