package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	fmt.Println("Current running go routines: ", runtime.NumGoroutine())
	fmt.Println("Running Sync")
	syncFunction()
	time.Sleep(1 * time.Second)
	fmt.Println("Running Async")
	asyncFunction()
	
}

func asyncFunction() {
	wg.Add(10)
	timeStart := time.Now()
	go sliceIterationAsync()
	go sliceIterationAsync()
	
	go sliceIterationAsync()
	go sliceIterationAsync()
	
	go sliceIterationAsync()
	go sliceIterationAsync()
	
	go sliceIterationAsync()
	go sliceIterationAsync()

	go sliceIterationAsync()
	go sliceIterationAsync()
	fmt.Println("Current running go routines: ", runtime.NumGoroutine())
	defer takeTotalTime(&timeStart)
	wg.Wait()
}

func syncFunction() {
	
	timeStart := time.Now()
	
	sliceIteration()
	sliceIteration()
	
	sliceIteration()
	sliceIteration()
	
	sliceIteration()
	sliceIteration()

	sliceIteration()
	sliceIteration()
	
	sliceIteration()
	sliceIteration()

	timeEnd := time.Now()
	duration := timeEnd.Sub(timeStart)

	fmt.Printf("Synchronous function took %s\n", duration)
}

func sliceIteration() {
	// time.Sleep(100 * time.Millisecond)
	iterations := 20_000_000

	// slice of string maps
	type KV struct {
		value int
	}
	// NORMAL IMPLEMENTATION
	objects := make([]KV, 0, iterations)

	for i := 0; i < iterations; i++ {
		objects = append(objects, KV{value: 1})
	}
}
func sliceIterationAsync() {
	// time.Sleep(100 * time.Millisecond)
	iterations := 20_000_000

	// slice of string maps
	type KV struct {
		value int
	}
	// NORMAL IMPLEMENTATION
	objects := make([]KV, 0, iterations)

	for i := 0; i < iterations; i++ {
		objects = append(objects, KV{value: 1})
	}
	wg.Done()
}


func takeTotalTime(timeStart *time.Time) {
	timeEnd := time.Now()
	duration := timeEnd.Sub(*timeStart)
	fmt.Printf("Asynchronous function took %s\n", duration)
}