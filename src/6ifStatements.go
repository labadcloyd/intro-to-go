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

	/*
	* SWITCH STATEMENTS
	*/

	// basic switch statement
	fmt.Printf("\n")
	switch 3 {
	case 1, 5, 10:
		fmt.Println("Case is either 1, 5, or 10")
	case 2 | 3, 4, 9:
		fmt.Println("Case is either 2, 3, 4, or 9")
	default:
		fmt.Println("Case is neither 1, 3, 4, 5, 9, or 19")
	}

	// you can write the same if statement in a switch
	// where you can do operations
	switch i := 3 + 2; i  {
	case 1, 5, 10:
		fmt.Println("Case is either 1, 5, or 10")
	case 2 | 3, 4, 9:
		fmt.Println("Case is either 2, 3, 4, or 9")
	default:
		fmt.Println("Case is neither 1, 3, 4, 5, 9, or 19")
	}
	i :=  5
	switch {
	case i > 5:
		fmt.Println("Case is greater than 5")
	case i < 5:
		fmt.Println("Case is lesser than 5")
	default:
		fmt.Println("Case is neither greater or lesser than 5")
	}

	// you can also check the types
	var j interface{} = [3]int{}
	switch j.(type) {
	case int:
		fmt.Println("J is an integer")
	case float64: 
		fmt.Println("J is a float64")
	case [2]int: 
		fmt.Println("J is an array of 2 integers")
	case [3]int: 
		fmt.Println("J is an array of 3 integers")
	default:
		fmt.Println("J is another type")
	}

}
