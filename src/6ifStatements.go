package main

import "fmt"

func main() {
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
		fmt.Printf("Florida Exists")
	}
}