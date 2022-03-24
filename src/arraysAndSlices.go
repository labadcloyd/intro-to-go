package main

import "fmt"

func main() {
	/*
		*ARRAYS
	*/
	grades := [3]int{87, 89, 88}
	fmt.Printf("\n%v, %T", grades, grades)

	// you dont have to specify the size when initiating an array
	grades2 := [...]int{87, 89, 88}
	fmt.Printf("\n%v, %T", grades2, grades2)

	// another way to declare arrays
	var students [4]string
	students[0] = "Cloyd"
	fmt.Printf("\nStudents: %v, %T", students, students)
	fmt.Printf("\nLength of array of students: %v, %T", len(students), students)

	// you can also declare multi dimensional arrays
	// using identity matrix as an example, this is related to linear algebra
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int { 1, 0, 0 }
	identityMatrix[1] = [3]int { 0, 1, 0 }
	identityMatrix[2] = [3]int { 0, 0, 1 }
	fmt.Printf("\n%v, %T", identityMatrix, identityMatrix)

	// when you try to copy an array you make an entirely new array
	// which apparently other languages don't function like that

	a := [6]int {6, 5, 4, 3, 2, 1} 
	b := a
	b[1] = 26
	fmt.Printf("\n%v", a)
	fmt.Printf("\n%v", b)

	f := a[2:] // array after 2nd element to end
	g := a[:4] // first 4 elements
	h := a[1:4] // 2nd element to 4th element
	fmt.Printf("\nArray after 2nd element to end %v", f)
	fmt.Printf("\nFirst 4 elements: %v", g)
	fmt.Printf("\n2nd element to 4th element: %v", h)

	/*
	* SLICES
	*/
	// slices are almost identical to arrays but with a few exceptions
	// slices dont have a specification on their length without us-
	// ing the three dots which we use with arrays
	c := []int {99, 98, 97}
	fmt.Printf("\nSlices: %v", c)

	// slices are naturally reference types
	// so when you copy an array to a variable, youre not creating a new array
	// instead youre making a refference
	d := c
	d[1] = 90
	fmt.Printf("\nOriginal Slices: %v", c)
	fmt.Printf("\nReference Slices: %v", d)

}