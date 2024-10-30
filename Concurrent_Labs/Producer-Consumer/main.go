package main

import (
	"fmt"
	"sync"
	"time"
)


func producer(mute *sync.Mutex, ch chan int) {
	fmt.Println("The producer is executing.")
	// here i iterate five times 
	for i:=0; i<5; i++ {
		mute.Lock()
		// here i put the value of i into the chanel named ch
		ch <- i
		// here i sleep for a milli-second
		time.Sleep(time.Millisecond)
		mute.Unlock()
	}
	fmt.Println("The producer is done executing.")
}


func consumer(ch chan int) {
	// this for loop gets the value(s) that are put into the chanel 
	for values := range <-ch {
		fmt.Println("Comsumer has: ", values)
	}
}


func main() {
	// here i get the current local time 
	now := time.Now()

	// i print how long it takes to run the program at the end of the main function
	defer func() {
		fmt.Println("time it took to execute producer / consumer task:", time.Since(now))
	}()

	var mute sync.Mutex
	ch := make(chan int, 5)
	threads := 20
	
	for i:=0; i<threads; i++ {
		go producer(&mute, ch)
		consumer(ch)
	}
}