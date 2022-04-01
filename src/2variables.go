package main

import (
	"fmt"
	"strconv"
)

var (
	firstName = "cloyd"
	lastName = "abad"
	age = 20
)

// Upper cased variables are accessible outside of the package or file
var I = 20
// Lower cased variables are only accessible inside the file

func main() {
	fmt.Println(firstName, lastName, age)
	
	// when using ":=" you no longer need the var keyword
	// ":=" is only used when declaring and assigning a new variable 
	firstName2 := "cloyd"
	lastName2 := "chan"
	age2 := 18

	fmt.Println(firstName2, lastName2, age2)

	var (
		i int
		j int
	)
	i = 100
	fmt.Println(i)

	j = 40.0
	fmt.Println(j)

	//converting types
	var floatingNumber float32 = 25.5
	var intNumber = 10
	intNumber = int(floatingNumber)

	fmt.Printf("%v, %T", intNumber, intNumber)

	//converting string
	var age = 20
	var ageString = "20"
	var stringAge = strconv.Itoa(age)
	fmt.Printf("\n %v, %T", stringAge, stringAge)
	
	newAge, err:= strconv.Atoi(ageString)
	fmt.Printf("\n %v, %T error: %v", newAge, newAge, err)

}