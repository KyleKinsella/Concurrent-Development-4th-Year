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

// here i have made a function to reduce the duplicated code that i did have in my code
func do(goNum int, sem *semaphore.Weighted, ctx context.Context, m *sync.Mutex) {
	m.Lock()
	sem.Acquire(ctx, 1)
	goNum++
	sem.Release(1)	
	defer m.Unlock()
}

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, wg *sync.WaitGroup, m *sync.Mutex, sem *semaphore.Weighted, ctx context.Context, max int) bool {
	
	// i make a new variable and assign it to the parameters of max
	maxVal := max
	fmt.Println("Part A",goNum) // all part a's first then all of part b's 

	//we wait here until everyone has completed part A
	time.Sleep(time.Second)
	
	// then i loop through the number of tasks to complete 
	for i :=0; i<maxVal; i++ {
		// then i run my do function with the go key word
		go do(goNum, sem, ctx, m)
	}

	fmt.Println("PartB",goNum)
	wg.Done()
	return true
}


func main() {	
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
	
	for i := range totalRoutines {//create the go Routines here
		// here i have added more parameters because i have updated my function called "doStuff"
		go doStuff(i, &wg, &theLock, sem, ctx, 11)
	}
	wg.Wait() //wait for everyone to finish before exiting
}
