package main

import "fmt"

func main() {
	// sayMessage("This message is from a function", 10)

	//* USING POINTERS IN FUNCTIONS
	/// if you want to manipulate data you usually use a function
	/// then return the data, it works but there's a more efficient way
	/// by passing in pointers you would no longer need to return data 
	/// but you can manipulate the data directly like so:
	// greeting := "Hi there,"
	// name := "Cloyd"
	// sayGreeting(greeting, &name)
	// fmt.Printf("updated: %v %v\n", greeting, name)

	//* DOT NOTATION IN FUNCTIONS
	/// you can join a group of numbers into a slice using the dot notation
	/// or also called as VARIADIC Parameters
	// numbers := []int{1,2,3,4,5,6}
	// joinSlice(numbers...)
	/// you can only have one VARIADIC Parameter and it has to be at the end
	joinSlice("The sum of all numbers is", 1, 2, 3, 4, 5, 6)

	//* RETURNING AN ERROR FROM A FUNCTION
	result, err := divide(5.0, 1.5)
	fmt.Println("result:", result, "error:", err)

	//* ANONYMOUS FUNCTION
	func(){
		fmt.Println("Im an anonymous function.")
	}()
}

func sayMessage(msg string, len int) {
	for i := 0; i < len; i++ {
		fmt.Printf("message: %v, index: %v\n", msg, i)
	}
}

func sayGreeting(greeting string, name *string) {
	fmt.Printf("%v %v\n", greeting, *name)
	*name = "chloe"
}

func joinSlice(msg string, numbers ...int){
	fmt.Println(numbers)
	result := 0
	for _, v := range numbers {
		result += v
	}
	fmt.Println(msg, result)
}

func divide (a float64, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot be divided by zero")
	}
	return a/b, nil
}