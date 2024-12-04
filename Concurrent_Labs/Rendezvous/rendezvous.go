package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	// here I define all of my variables that I will be using
	globalMute sync.Mutex
	wait sync.WaitGroup
	counter int
	cond = sync.NewCond(&globalMute)
)

func runRendezvous() {

	one := func() {
		fmt.Println("THREAD 1: Starting part one")
		time.Sleep(time.Second)
		// wait

		// here I lock the mutex to stop any other thread getting access 
		globalMute.Lock()
		
		// counter is how many threads have arrived 
		counter++
		if counter < 2 {
			// wait until we have 2 threads
			cond.Wait()
		} else {
			// when counter reaches 2, we wake up all waiting goroutines
			cond.Broadcast()
		}

		// unlock the mutex to allow other threads to gain access 
		globalMute.Unlock()
		fmt.Println("THREAD 1: Starting part two")

		// decrement the waitgroup by one 
		wait.Done()
	}

	two := func() {
		fmt.Println("THREAD 2: Starting part one")
		time.Sleep(time.Second)
		// wait
		
		// here I lock the mutex to stop any other thread getting access 
		globalMute.Lock()

		// counter is how many threads have arrived 
		counter++
		if counter < 2 {
			// wait until we have 2 threads
			cond.Wait()
		} else {
			// when counter reaches 2, we wake up all waiting goroutines
			cond.Broadcast()
		}

		//  unlock the mutex to allow other threads to gain access 
		globalMute.Unlock()
		fmt.Println("THREAD 2: Starting part two")
		
		// decrement the waitgroup by one 
		wait.Done()
	}
	// wait for 2 goroutines to run
	wait.Add(2)
	// run the one and two functions
	go one()
	go two()
	// wait until the waitgroup hits zero
	wait.Wait()
} 

func main() {
	// run my function
	runRendezvous()
}
