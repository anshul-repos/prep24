package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID int
}

type WorkerPool struct {
	numberWorker int
	jobQueue     chan Job
	results      chan int
	wg           sync.WaitGroup
}

func NewWorkerPool(nWorker, nJob int) *WorkerPool {
	return &WorkerPool{
		numberWorker: nWorker,
		jobQueue:     make(chan Job, nJob),
		results:      make(chan int, nJob),
	}
}

func (wp *WorkerPool) AddJob(nJob int) {
	for i := 1; i <= nJob; i++ {
		wp.jobQueue <- Job{ID: i}
	}
}

func (wp *WorkerPool) StartWorker() {
	for i := 1; i <= wp.numberWorker; i++ {
		wp.wg.Add(1)
		go wp.Worker(i)
	}
}

func (wp *WorkerPool) Worker(ID int) {
	defer wp.wg.Done()
	for job := range wp.jobQueue {
		fmt.Printf("Worker %d started Job(%d)\n", ID, job.ID)
		time.Sleep(time.Second * 2)
		fmt.Printf("Worker %d finished Job(%d)\n", ID, job.ID)
		wp.results <- job.ID
	}
}

func (wp *WorkerPool) CollectResults() {
	for res := range wp.results {
		fmt.Printf("Recieved result for job(%d) \n", res)
	}
}

func main() {
	noOfWorker, noOfJobs := 3, 10
	wp := NewWorkerPool(noOfWorker, noOfJobs)
	wp.AddJob(noOfJobs)
	close(wp.jobQueue)
	wp.StartWorker()
	wp.wg.Wait()
	close(wp.results)
	wp.CollectResults()
}
