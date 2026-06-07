/*
You are building a backend service for a distributed scheduling system.
You need to provide a service that allows other parts of the application
to 'submit' tasks to be executed after a specific delay
(e.g., send an email 10 seconds after a user registers).

Your requirements are:
Non-blocking: The Submit call must return immediately, even if the delay is 1 hour.
Scalable: The system must handle thousands of concurrent tasks without crashing the host.
No time.Sleep in the main flow: You cannot block the caller by sleeping.
Graceful: If the service is shutting down, it should ideally stop processing pending tasks.
*/

package main

import (
	"fmt"
	"time"
)

type Job struct {
	RunAt time.Time
	Fn    func()
}

type Scheduler struct {
	// What fields do you need here?
	// Think about how to keep track of the tasks or how to dispatch them.
}

func NewScheduler() *Scheduler {
	return &Scheduler{}
}

func (s *Scheduler) Submit(task func(), delay time.Duration) {
	// Implement the non-blocking logic here.

	go func() {
		time.Sleep(delay)
		task()
	}()

}

func main() {
	s := NewScheduler()
	s.Submit(func() {
		fmt.Printf("I was executed!")
	}, 10*time.Millisecond)

	time.Sleep(100 * time.Millisecond)
}
