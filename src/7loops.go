package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v, ", i)
	}
	fmt.Printf("\n")

	// you can also declare the variable outside of the loop
	// so you can use it
	j := 0
	for ; j < 10; j++ {
		fmt.Printf("%v, ", j)
	}
	fmt.Printf("\n")
	fmt.Printf("%v", j)
	fmt.Printf("\n")

	// there is no "while" keyword in go
	// instead you can use the "for" keyword to make a while loop
	j = 0
	for j < 10 {
		fmt.Printf("%v, ", j)
		j++
	}
	fmt.Printf("\n")

	// you can also skip an iteration of a loop using the continue keyword
	j = 0
	for j < 10 {
		// only printing the even data
		if j % 2 == 0 {
			fmt.Printf("%v, ", j)
			j++
		} else {
			j++
			continue
		}
	}
	fmt.Printf("\n")

	// when using a nested loop the break keyword will only break the loop
	// that is where it's nearest to
	fmt.Printf("\nThe break keyword not working as intended\n")
	for i:= 1; i < 5; i++ {
		for j = 1; j < 5; j++ {
			fmt.Printf("%v, ", i*j)
			if i*j >= 5 {
				break
			}
		}
	}
	fmt.Printf("\n")
	// a solution to this is by using labels
	// we then use the label to identify what loop to break
	fmt.Printf("\nThe break keyword now working\n")
	loop:
		for i:= 1; i < 5; i++ {
			for j = 1; j < 5; j++ {
				if i*j >= 5 {
					break loop
				}
				fmt.Printf("%v, ", i*j)
			}
		}
	fmt.Printf("\n")
	fmt.Printf("\nOriginal values before the using the break keyword\n")
	for i:= 1; i < 5; i++ {
		for j = 1; j < 5; j++ {
			fmt.Printf("%v, ", i*j)
		}
	}
	fmt.Printf("\n")

	// looping through a slice
	sample := make([]int, 0, 100)
	sample = append(sample, 24, 25, 26, 27, 28, 29, 30)
	fmt.Printf("\nLooping through a slice\n")
	for j = 0; j < len(sample); j++ {
		fmt.Printf("%v, ", sample[j])
	}
	fmt.Printf("\n")
	// a declarative way of looping through a slice
	// would be by using range
	sample2 := make([]int, 0, 100)
	sample2 = append(sample2, 24, 25, 26, 27, 28, 29, 30)
	fmt.Printf("\nLooping through a slice declaratively\n")
	for key, value := range sample2 {
		fmt.Printf("index: %v, value: %v\n", key, value)
	}
	fmt.Printf("\n")
	// you can also loop through a string declaritively
	sample3 := "Hello World"
	fmt.Printf("\nLooping through a slice declaratively\n")
	for key, value := range sample3 {
		fmt.Printf("string:%v, index: %v, unicode-value: %v\n", string(value), key, value)
	}
	fmt.Printf("\n")
	

}