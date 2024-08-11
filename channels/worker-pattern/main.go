// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func worker(id int, jobs <-chan int, results chan<- int) {
// 	for job := range jobs {
// 		fmt.Printf("Worker %d processing job %d\n", id, job)

// 		results <- job * 2
// 	}

// }
// func main() {

// 	jobs := 10
// 	workers := 3

// 	send := make(chan int, jobs)
// 	results := make(chan int, jobs)

// 	wg := sync.WaitGroup{}

// 	for i := 1; i <= workers; i++ {
// 		wg.Add(1)

// 		go func(workerId int) {
// 			defer wg.Done()
// 			worker(workerId, send, results)

// 		}(i)
// 	}
// 	go func() {
// 		wg.Add(1)
// 		defer wg.Done()

// 		for i := 1; i <= jobs; i++ {
// 			send <- i
// 		}
// 		close(send)

// 	}()

// 	// // Wait for all workers to finish
// 	go func() {

// 		wg.Wait()
// 		close(results)

// 	}()

// 	for result := range results {
// 		fmt.Printf("Result: %d\n", result)

// 	}

// }

package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		results <- job * 2
	}
	time.Sleep(time.Second)

}

func main() {
	numJobs := 10
	numWorkers := 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	status := make(chan bool)

	// Start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		go func(workerID int) {
			worker(workerID, jobs, results)
			status <- true
		}(i)
	}

	// Enqueue jobs
	go func() {
		for i := 1; i <= numJobs; i++ {
			jobs <- i
		}
		status <- true
		close(jobs)

	}()

	go func() {
		<-status
		<-status
		close(results)

	}()

	// Collect results
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}

}
