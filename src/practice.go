package main

import (
	"fmt"
	"time"
)

func main() {

	iterations := 10_000_000

	// slice of string maps
	type KV struct {
		value string
	}
	// A FASTER IMPLEMENTATION
	objects1 := make([]KV, iterations)
	timeStart1 := time.Now()

	for i := range objects1 {
		objects1[i] = KV{value: "name"}
	}

	timeEnd1 := time.Now()
	duration1 := timeEnd1.Sub(timeStart1)

	fmt.Println("len:", len(objects1))
	fmt.Printf("%d iterations took %s\n", iterations, duration1)

	// NORMAL IMPLEMENTATION
	objects := make([]KV, 0, iterations)
	timeStart := time.Now()

	for i := 0; i < iterations; i++ {
		objects = append(objects, KV{value: "name"})
	}

	timeEnd := time.Now()
	duration := timeEnd.Sub(timeStart)

	fmt.Println("len:", len(objects))
	fmt.Printf("%d iterations took %s\n", iterations, duration)

	// slice of string maps
	timeStartAllocated := time.Now()
	objectsAllocated := make([]KV, iterations, iterations)

	for i := 0; i < iterations; i++ {
		objectsAllocated[i] = KV{value: "name"}

	}

	timeEndAllocated := time.Now()
	durationAllocated := timeEndAllocated.Sub(timeStartAllocated)

	fmt.Println("len:", len(objectsAllocated))
	fmt.Printf("%d iterations took (with pre-allocation) %s\n", iterations, durationAllocated)

}