package main

import (
	"fmt"
)

func main() {
	const a int = 2
	fmt.Printf("%v, %T\n", a, a)

	// !you cannot set constants with a value that has to be determined at run time
	// const b float64 = math.Sin(1.57)

	// you can also add constants as long as theyre the same type
	const b int = 5
	var c int = 10
	fmt.Printf("%v, %T\n", b+c, b+c)
	
	// you can also do implicit conversions with constants when using integers
	const d = 20
	var e int16 = 10
	fmt.Printf("%v, %T\n", d+e, d+e)
	// !but cannot be don with other types like floats
	// const d = 20.5
	// var e int16 = 10
	// fmt.Printf("%v, %T\n", d+e, d+e)
}