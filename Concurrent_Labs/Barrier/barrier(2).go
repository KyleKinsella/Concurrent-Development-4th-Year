
//Barrier.go Template Code
//Copyright (C) 2024 Dr. Joseph Kehoe

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.


//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Kyle Kinsella - C00273146
// Issues:
// The barrier is not implemented! / I have now implemented a barrier with a mutex added to the function called "doStuff"
//--------------------------------------------

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
	"golang.org/x/sync/semaphore"
)


// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, wg *sync.WaitGroup, m *sync.Mutex) bool {
	// time.Sleep(time.Second)
	fmt.Println("Part A",goNum) // all part a's first then all of part b's 
	//we wait here until everyone has completed part A


	// here i am locking my mutex lock
	m.Lock()
	goNum++
	// ch <- goNum
	// here i sleep for one second, this allows the a's to be printed then after one second the b's print
	time.Sleep(time.Second)
	// here i unlock my mutex
	m.Unlock()
	

	fmt.Println("PartB",goNum)
	wg.Done()
	return true
}


func main() {
	// ch := make(chan int)

	
	// here i get the current local time 
	now := time.Now()

	// i print how long it takes to run the program at the end of the main function
	defer func() {
		fmt.Println(time.Since(now))
	}()

	totalRoutines:=10
	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	//we will need some of these
	ctx := context.TODO()
	var theLock sync.Mutex
	sem := semaphore.NewWeighted(int64(totalRoutines))
	theLock.Lock()
	sem.Acquire(ctx, 1)

	// i define a mutex lock to pass into my function
	var mute sync.Mutex

	for i := range totalRoutines {//create the go Routines here
		// i pass in a refernce my mutex using a & symbol
		// ch <- i
		go doStuff(i, &wg, &mute)
		// <- ch
	}
	sem.Release(1)
	theLock.Unlock()
	
	wg.Wait() //wait for everyone to finish before exiting
}
