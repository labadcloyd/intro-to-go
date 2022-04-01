package main

import "fmt"

func main() {
	/*
	* IF STATEMENTS
	*/
	if true {
		fmt.Printf("True")
		fmt.Printf("\n")
	}
	// you can create a conditional if you want to check-
	// if a certain field has a value
	var statePopulation = map[string]int{
		"California": 391000,
		"Illinois":   251000,
		"Washington": 281000,
		"Nevada":     211000,
		"Florida":    199000,
	}

	if _, floridaExists := statePopulation["Florida"]; floridaExists {
		fmt.Printf("Florida Exists\n")
	}

	// basic conditional operators
	num1 := 30
	num2 := 40

	fmt.Println(num1 <= num2, num1 >= num2, num1 != num2)
	fmt.Println(num1 <= num2 || num1 >= num2)
	fmt.Println(num1 <= num2 && num1 >= num2)
	fmt.Println(num1 <= num2 || !(num1 >= num2))
	fmt.Println(num1 <= num2 && !(num1 >= num2))

	// when writing conditionals with logical operators
	// go checks it from left to right, and in an or operator
	// if one is true then go will no longer check the remaining
	// statements
	fmt.Printf("\n")
	returnTrue := func() bool {
		// functions cannot be declared inside a function unless
		// you use variables and assign a function to it
		fmt.Println("Returned True")
		return true
	}

	if num1 < 1 || returnTrue() || num1 > 1 {
		fmt.Println("function return true was called")
	}
	if num1 > 1 || returnTrue() || num1 > 1 {
		fmt.Println("function return true was NOT called")
		fmt.Printf("\n")
	}
	if num1 > 1 && returnTrue() {
		fmt.Println("function return true was called")
	}
	if num1 < 1 && returnTrue() {
		fmt.Println("function return true was NOT called")
	}



}
