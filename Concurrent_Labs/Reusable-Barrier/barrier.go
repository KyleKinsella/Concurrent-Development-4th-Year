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
// The barrier is not implemented! / I have now implemented a barrier with a mutex and semaphores added to the function called "doStuff"
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
func doStuff(goNum int, wg *sync.WaitGroup, m *sync.Mutex, sem *semaphore.Weighted, ctx context.Context) bool {//, sem *semaphore.Weighted, ctx context.Context, max int) bool {

	// here I am defering the waitgroup until the end of the method is called 
	defer wg.Done()

	fmt.Println("Part A",goNum) // all part a's first then all of part b's 

	// here I am looping 2 times to print, a's followed by each iteration then then b's followed by each iteration 
	for i:=0; i<2; i++ {
		// here I lock my mutex
		m.Lock()
		// here I aquire a context and a value of 1 
		sem.Acquire(ctx, 1)
		// here I increment goNum
		goNum++
		fmt.Println("Iteration: ", i)
		time.Sleep(time.Second)
		// here I unlock my mutex
		m.Unlock()
		// then I relase the value of 1 that I previously aquired
		sem.Release(1)
	}
	fmt.Println("PartB:", goNum)
	return true
}

func main() {	
	// here i get the current local time 
	now := time.Now()

	// i print how long it takes to run the program at the end of the main function
	defer func() {
		fmt.Println(time.Since(now))
	}()

	totalRoutines := 2
	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	ctx := context.TODO()
	var theLock sync.Mutex
	sem := semaphore.NewWeighted(int64(totalRoutines))

	for i := range totalRoutines {//create the go Routines here
		// here i have added more parameters because i have updated my function called "doStuff"
		go doStuff(i, &wg, &theLock, sem, ctx)
	}
	wg.Wait() //wait for everyone to finish before exiting		
}
