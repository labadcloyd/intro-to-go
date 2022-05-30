package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	///* CHANNELS
	// using channels is a way to pass data accross different go routines.

	// we can create a channel through the make method
	// then we will specify the data that will pass through the channel
	// the second argument will be the capacity of the channel
	ch := make(chan int, 2)

	wg.Add(4)
	go sendValue(ch, 47)
	go sendValue(ch, 27)
	go recieveValue(ch)
	go recieveValue(ch)

	wg.Wait()
}

func sendValue(c chan int, value int) {
	defer wg.Done()
	fmt.Println("Processing go routine", value)
	c <- value
	fmt.Println("Finished processing go routine")
}
func recieveValue(c chan int) {
	defer wg.Done()
	value := <- c
	fmt.Println(value)
}