package main

import "fmt"

func main() {
	/*
		*ARRAYS
	*/
	grades := [3]int{87, 89, 88}
	fmt.Printf("%v, %T", grades, grades)

	// you dont have to specify the size when initiating an array
	grades2 := [...]int{87, 89, 88}
	fmt.Printf("%v, %T", grades2, grades2)

	// another way to declare arrays
	var students [4]string
	students[0] = "Cloyd"
	fmt.Printf("Students: %v, %T", students, students)
	fmt.Printf("Length of array of students: %v, %T", len(students), students)

	
}