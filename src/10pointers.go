package main

import "fmt"

func main() {
	/*
	*	POINTERS
	*/
	var a int = 42
	var b int = a
	a = 50
	fmt.Printf("POINTERS\n")
	fmt.Printf("Variables without pointers:\n")
	fmt.Printf("a: %v, b: %v\n", a, b)
	var c int = 42
	var d *int = &c
	c = 50
	fmt.Printf("Variables with pointers:\n")
	fmt.Printf("c: %v, b: %v\n", c, *d)
	// in order to update a pointer variable
	// you derefference it (which gets the value from that address)-
	// and update it
	*d = 32
	fmt.Printf("Variables with pointers:\n")
	fmt.Printf("c: %v, b: %v\n", c, *d)


	
	fmt.Printf("\nHOW POINTERS' ADDRESSES ARE SAVED\n")
	q := [4]int{1,2,3,4}
	w := &q[2]
	e := &q[3]
	fmt.Printf("Addresses of pointers:\n")
	fmt.Printf("q: %v, w: %p, e: %p\n", q, w, e)
	// notice the difference of the addresses
	// they have a gap of 8 since an integer in go is 8 bytes long




	// using pointers for structs
	fmt.Printf("\nUSING POINTERS FOR STRUCTS\n")
	type myStruct struct{
		foo int
	}
	var ms *myStruct
	ms = &myStruct{foo: 20}
	fmt.Println(ms)
	
}