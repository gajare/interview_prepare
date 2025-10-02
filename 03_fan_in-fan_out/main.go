package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := make(chan int, 10)

	var wg sync.WaitGroup
	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go Worker(i, jobs, &wg)
	}

	for j := 1; j <= 1000; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()
	fmt.Println("All jobs processed...!")
}

func Worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		//time.Sleep(10 * time.Millisecond)
		fmt.Println("worker id :", id, "process job:", job)
	}
}
