package main

import (
	"sync"
)

func worker(id int, job <- chan int, wg *sync.WaitGroup){
	defer wg.Done()
	for j := range job {
		println("worker", id, "started job", j)
	}	
}

func main() {
	jobs := make(chan int, 100)
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i,jobs, &wg)
	}

	for j := 1; j <= 10; j++{
		jobs <- j;
	}

	close(jobs)
	wg.Wait()
}