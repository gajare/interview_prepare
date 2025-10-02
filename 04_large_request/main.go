package main

import (
	"fmt"
	"net/http"
	"time"
)

const workerCount = 10
const jobQueueSize = 1000

var jobQueue = make(chan int, jobQueueSize)

func worker(id int, jobs <-chan int) {
	for j := range jobs {
		fmt.Printf("worker %d processing job %d\n", id, j)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	for i := 0; i <= workerCount; i++ {
		go worker(i, jobQueue)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		select {
		case jobQueue <- time.Now().Nanosecond():
			fmt.Fprintf(w, "job accepted\n")
		default:
			http.Error(w, "Too many request try later", http.StatusTooManyRequests)
		}
	})
	http.ListenAndServe(":8080", nil)
}
