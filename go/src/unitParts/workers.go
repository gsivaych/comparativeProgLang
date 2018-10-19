package main

import (
	"fmt"
)

type Worker interface {
	Run() interface{}
}

type WorkQueue struct {
	Jobs         chan Worker
	Results      chan interface{}
	StopRequests chan int
	NumWorkers   uint
}

// Create a new work queue capable of doing nWorkers simultaneous tasks, expecting to queue maxJobs tasks.
func Create(nWorkers uint, maxJobs uint) *WorkQueue {
	q := new(WorkQueue)
	// initialize struct; start nWorkers workers as goroutines
	q.Jobs = make(chan Worker)
	q.Results = make(chan interface{})
	q.StopRequests = make(chan int)
	q.NumWorkers = nWorkers
	for i := 0;i < int(nWorkers);i++ {
		go q.worker()
	}
	return q
}

// A worker goroutine that processes tasks from .Jobs unless .StopRequests has a message saying to halt now.
func (queue WorkQueue) worker() {
	running := true
	// Run tasks from the Jobs channel, unless we have been asked to stop.
	for running {
		for {
        select {
        case signal := <- queue.StopRequests:
						if signal > 0 {
							running = false
						}
            break
        default:
        }
        select {
        case signal := <- queue.StopRequests:
						if signal > 0 {
							running = false
						}
            break
        case job := <- queue.Jobs:
						val := job.Run()
						queue.Results <- val
        }
        break
    }
	}
}

func (queue WorkQueue) Enqueue(work Worker) {
	// put the work into the Jobs channel so a worker can find it and start the task.
	queue.Jobs <- work
}

func (queue WorkQueue) Shutdown() {
	// tell workers to stop processing tasks.
	for i := 0;i < int(queue.NumWorkers);i++ {
		queue.StopRequests <- 1
	}
}


func main(){
	fmt.Println("ting !!")
	qu := Create(2,3)
	qu.StopRequests <- 1
	// qu.Shutdown()
	fmt.Println("ting !!")
}
