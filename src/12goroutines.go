package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	///* GO ROUTINES
	// go routines are executed outside of the main function
	// so even if the go routine isn't finished executing but
	// the main function is done, then the print statement in 
	// this sample will not work

	// however a way around this would be by using wait groups
	// we start off by adding a wait group, in this case just 1
	wg.Add(1)
	msg := "Hello World"
	go func() {
		fmt.Println(msg)
	//	the done method then decrements the number of wait groups
		wg.Done()
	}()
	
	// we can also spawn multiple go routines
	// the problem with wait groups is that go routines
	// are not executed in order
	// a way to fix this is by using a mutix
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go greet("Hello World")
		m.Lock()
		go increment()
	}
	// however this implementation forces the process to be done synchrnously
	// but there may be situations where you may use this.


	// finally the wait group will wait until all the wait groups are finished
	wg.Wait()
}

func greet(msg string) {
	fmt.Println(msg, counter)
	m.RUnlock()
	wg.Done()
}
func increment() {
	counter++
	m.Unlock()
	wg.Done()
}