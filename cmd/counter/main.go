package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type mutexCounter struct {
	mu    sync.Mutex
	value int
}

func (c *mutexCounter) Inc() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

type atomicCounter struct {
	value int64
}

func (c *atomicCounter) Inc() {
	atomic.AddInt64(&c.value, 1)
}

func main() {
	const numGoroutines = 1000
	const numIncrementsPerGoroutine = 10000

	//Test 1: Mutex Counter
	mc := &mutexCounter{}
	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < numIncrementsPerGoroutine; j++ {
				mc.Inc()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	mutexTime := time.Since(start)

	//Test 2: Atomic Counter
	ac := &atomicCounter{}
	start = time.Now()

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < numIncrementsPerGoroutine; j++ {
				ac.Inc()
			}
			wg.Done()
		}()
	}

	wg.Wait()

	atomicCounterTime := time.Since(start)

	fmt.Printf("Mutex: %d increments in %v\n", mc.value, mutexTime)
	fmt.Printf("Atomic: %d increments in %v\n", ac.value, atomicCounterTime)
	fmt.Printf("Atomic is %.1fx faster than mutex\n", float64(mutexTime)/float64(atomicCounterTime))
}
