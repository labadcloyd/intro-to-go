package main

import "fmt"

func main() {
	/*
		*ARRAYS
	*/
	grades := [3]int{87, 89, 88}
	fmt.Printf("\n%v, %T", grades, grades)
	fmt.Printf("\n")


	// you dont have to specify the size when initiating an array
	grades2 := [...]int{87, 89, 88}
	fmt.Printf("\n%v, %T", grades2, grades2)
	fmt.Printf("\n")


	// another way to declare arrays
	var students [4]string
	students[0] = "Cloyd"
	fmt.Printf("\nStudents: %v, %T", students, students)
	fmt.Printf("\nLength of array of students: %v, %T", len(students), students)
	fmt.Printf("\n")


	// you can also declare multi dimensional arrays
	// using identity matrix as an example, this is related to linear algebra
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int { 1, 0, 0 }
	identityMatrix[1] = [3]int { 0, 1, 0 }
	identityMatrix[2] = [3]int { 0, 0, 1 }
	fmt.Printf("\n%v, %T", identityMatrix, identityMatrix)
	fmt.Printf("\n")


	// when you try to copy an array you make an entirely new array
	// which apparently other languages don't function like that

	a := [6]int {6, 5, 4, 3, 2, 1} 
	b := a
	b[1] = 26
	fmt.Printf("\n%v", a)
	fmt.Printf("\n%v", b)
	fmt.Printf("\n")

	f := a[2:] // elements skipped from the start
	g := a[:3] // first 3 elements
	h := a[1:4] // 2nd element to 4th element
	i := a[:] // all elements
	fmt.Printf("\nElements skipped from the start: %v", f)
	fmt.Printf("\nFirst 3 elements: %v", g)
	fmt.Printf("\n2nd element to 4th element: %v", h)
	fmt.Printf("\nAll Elements: %v", i)
	fmt.Printf("\n")

	/*
	* SLICES
	*/
	// slices are almost identical to arrays but with a few exceptions
	// slices dont have a specification on their length without us-
	// ing the three dots which we use with arrays
	c := []int {99, 98, 97}
	fmt.Printf("\nSlices: %v", c)
	fmt.Printf("\n")

	// slices are naturally reference types
	// so when you copy a slice to a variable, youre not creating a new slice
	// instead youre making a refference
	d := c
	d[1] = 90
	fmt.Printf("\nOriginal Slices: %v", c)
	fmt.Printf("\nReference Slices: %v", d)
	fmt.Printf("\n")
	
	// when creating slices you can increase the capacity by appending data
	a2 := []int{}
	fmt.Printf("\nLength: %v, Capacity: %v", len(a2), cap(a2))
	fmt.Printf("\nvalue: %v ", a2)
	a2 = append(a2, 3, 4, 5, 6)
	fmt.Printf("\nLength: %v, Capacity: %v", len(a2), cap(a2))
	fmt.Printf("\nvalue: %v ", a2)
	fmt.Printf("\n")
	// however using this way can be expensive when dealing with larger slices
	// an alternative to this would by using the make function on declaring the var
	// because this way the compiler can know ahead of time the capacity of the slice
	a3 := make([]int, 0, 100) 
	fmt.Printf("\nDeclaring slices using the make function")
	fmt.Printf("\nLength: %v, Capacity: %v", len(a3), cap(a3))
	fmt.Printf("\nvalue: %v ", a3)
	a3 = append(a3, 3, 4, 5, 6)
	fmt.Printf("\nLength: %v, Capacity: %v", len(a3), cap(a3))
	fmt.Printf("\nvalue: %v ", a3)
	fmt.Printf("\n")
	
	// go also has the equivalent of javascript's spread operator
	a3 = append(a3, []int{ 3, 4, 5, 6 }...)
	fmt.Printf("\nUsing the spread operator")
	fmt.Printf("\nLength: %v, Capacity: %v", len(a3), cap(a3))
	fmt.Printf("\nvalue: %v ", a3)
	fmt.Printf("\n")
	
	// how to pop elements off a slice
	b2 := a3[1 : len(a3)-1]
	fmt.Printf("\nPopping off elements from a slice")
	fmt.Printf("\nLength: %v, Capacity: %v", len(b2), cap(b2))
	fmt.Printf("\nvalue: %v ", b2)
	fmt.Printf("\n")

	// popping off an element in the middle is tricky
	// but you can pull it off by appending
	// first param is taking the first 3 elements
	// second param is the elements after the first 5 elements
	b2 = append(a3[:3], a3[5:]...)
	fmt.Printf("\nPopping off middle elements from a slice")
	fmt.Printf("\nLength: %v, Capacity: %v", len(b2), cap(b2))
	fmt.Printf("\nvalue: %v ", b2)
	//! becareful when using this as it will make changes to-
	//! the original slice
	fmt.Printf("\nOriginal Slice")
	fmt.Printf("\nLength: %v, Capacity: %v", len(a3), cap(a3))
	fmt.Printf("\nvalue: %v ", a3)
	fmt.Printf("\n")

}