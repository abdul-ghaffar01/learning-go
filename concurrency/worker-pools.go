package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID int
}

func worker(i int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {

		fmt.Printf("\nWorker %d Processing job %d...\n", i, job.ID)
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("Processed job %d\n", job.ID)

	}
}

func WorkerPool() {

	var wg sync.WaitGroup

	jobs := make(chan Job)

	workers := 5

	for i := range workers {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	for i := range 100 {
		jobs <- Job{ID: i}
	}

	close(jobs)

	fmt.Println("Program finished")
}
