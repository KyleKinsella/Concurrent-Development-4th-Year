package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand/v2"
	"testing"
)
//Global variables shared between functions --A BAD IDEA


func cs(status bool) {
	fmt.Println("Critical section")

	if(status == false) {
		fmt.Println("you are not in the critical section")
	}
	fmt.Println("you are in the CS.....")
}


var globalMute sync.Mutex
var wait sync.WaitGroup


func gr() {

	one := func() {
		fmt.Println("THREAD 1")
		fmt.Println("part one")
		// wait
		globalMute.Lock()
		fmt.Println("part two")
		wait.Done()
		globalMute.Unlock()
	}

	two := func() {
		
		fmt.Println("THREAD 2")
		fmt.Println("part one")
		wait.Wait()

		globalMute.Unlock()
		fmt.Println("part two")
		wait.Done()
		globalMute.Unlock()
	}

	wait.Add(2)
	go one()
	go two()
	wait.Wait()

} 


func check(n int) {

	var mutex1 sync.WaitGroup
	var mutex2 sync.WaitGroup


	// var mute sync.Mutex
	// var wg sync.WaitGroup

	// var count int

	for i:=0; i<n; i++ {
		fmt.Println("thread one is waiting")
		mutex1.Wait()
		// wg.Wait()
		globalMute.Lock()

		
		count := 0
		count++
		
		globalMute.Unlock()
		// mute.Lock()
		
		// cs area
		// fmt.Println("count variable has been updated")
		// count++
		// cs()
		// cs area
		
		// mute.Unlock()
		
		fmt.Println("thread two is waiting")
		
		
		wait.Wait()

		

		cs(true)

		defer mutex2.Done()
		// mute.Lock()
		// wg.Wait()
	}
} 



func WorkWithRendezvous(wg *sync.WaitGroup, Num int) bool {
	var X time.Duration
	X=time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second)//wait random time amount
	fmt.Println("Part A", Num)
	
	//Rendezvous here
	n := 2
	check(n)

	// gr()

	fmt.Println("PartB",Num)
	wg.Done()
	return true
}


func TestA(t *testing.T) {
	cs := "Critical section"
	fmt.Println(cs)
}




func main() {
	var wg sync.WaitGroup
	//barrier := make(chan bool)
	threadCount:=2

	wg.Add(threadCount)
	for N := range threadCount {
		go WorkWithRendezvous(&wg, N)
	}
	wg.Wait() //wait here until everyone (10 go routines) is done


	// da := <-time.After(X * time.Second)
	// fmt.Println("time to execute: ", ticker.Second())

	// ticker := time.NewTicker(1 * time.Second)
	// defer ticker.Stop()



	// for i:=0; i<5; i++ {
	// 	// Wait for the next tick
	// 	t := <-ticker.C
	// 	// Print only the seconds part
	// 	fmt.Println("Seconds:", t.Second())
	// 	}

	// <-ticker.C
}