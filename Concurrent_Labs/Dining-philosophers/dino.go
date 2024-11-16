//--------------------------------------------
// Author: Kyle Kinsella (C00273146@setu.ie)
// Created on 1/11/24
// Modified by: Kyle Kinsella - C00273146
// Issues:
// It can Deadlock! / it does not deadlock now.
//--------------------------------------------

package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

// here i added a method to remove duplicated code that was in think and eat
func removeDuplicateCode(index int) {
	var X time.Duration
	X=time.Duration(rand.Intn(5))
	time.Sleep(X * time.Second)//wait random time amount
}

func think(index int) {
	removeDuplicateCode(index)
	fmt.Println("Phil: ",index,"was thinking")
}

func eat(index int) {
	removeDuplicateCode(index)
	fmt.Println("Phil: ",index,"was eating")
}

// when a philosopher get a left fork and a right fork I use a mutex to lock it to that philosopher 
func getForks(leftFork, rightFork *sync.Mutex) {
	leftFork.Lock()
	rightFork.Lock()
}

// when a philosopher is done with the left fork and a right fork I unlock the mutex from that philosopher 
func putForks(leftFork, rightFork *sync.Mutex) {
	leftFork.Unlock()
	rightFork.Unlock()
}

func doPhilStuff(index int, wg *sync.WaitGroup, leftFork, rightFork *sync.Mutex){
	defer wg.Done()
	for i:=0; i<index; i++ {
		// here i call the think method
		think(index)
		// here i call the getForks method
		getForks(leftFork, rightFork)
		// here i call the eat method
		eat(index)
		// here i call the putForks method
		putForks(leftFork, rightFork)
	}
}

func main() {
	var wg sync.WaitGroup
	philCount := 5
	wg.Add(philCount)

	// here i make an array of mutex pointers with the size of 5
	forks := make([]*sync.Mutex, philCount)
	
	for i:=0; i<philCount; i++ {
		// each index in forks has a new mutex, it is made and put into forks[i]
		forks[i] = &sync.Mutex{} // after the =, creates a pointer to a new mutex
	
		// purpose: each philosopher has a fork represented by a mutex
	}

	for i:=0; i<philCount; i++ {
		// here the left fork is at forks[i]
		left := forks[i]
		// here the right fork is forks[i+1%philCount], this is for the wrap-around for the table
		right := forks[(i+1)%philCount]
		// the below if statement ensure that no deadlock occurs, because if true pick upthe right and left fork, otherwise false 
		// pick up the left and right fork. 
		if i == philCount-1 {
			// here we get the right fork then the left fork
			go doPhilStuff(i, &wg, right, left)
		} else {
			// here we get the left fork then the right fork
			go doPhilStuff(i, &wg, left, right)
		}
	}
	wg.Wait() // wait until everything is done
}
