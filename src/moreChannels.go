package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
func a() {
	defer wg.Done()
	for i := 0; i < 50; i++ {
		fmt.Printf("a")
	}
}
func b() {
	defer wg.Done()
	for i := 0; i < 50; i++ {
		fmt.Printf("b")
	}
}
func main() {
	wg.Add(4)
	go a()
	go b()
	go a()
	go b()
	wg.Wait()
	fmt.Printf("\nend main()")
}