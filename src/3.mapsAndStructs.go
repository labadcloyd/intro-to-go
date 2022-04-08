package main

import "fmt"

func main() {
	/*
	* MAPS
	*/
	//declaring a map
	// statePopulation := map[string]int{
	// 	"California": 391000,
	// 	"Illinois":   251000,
	// 	"Washington": 281000,
	// 	"Nevada":     211000,
	// 	"Florida":    199000,
	// }

	// another way to declare a map
	// this way we can dynamically assign value to a map
	statePopulation := make(map[string]int)
	statePopulation = map[string]int{
		"California": 391000,
		"Illinois":   251000,
		"Washington": 281000,
		"Nevada":     211000,
		"Florida":    199000,
	}

	fmt.Printf("Value: %v\n", statePopulation)
	fmt.Printf("Value: %v\n", statePopulation["Nevada"])
	fmt.Printf("\n")
	
	// we can also add values to maps
	statePopulation["Georgia"] = 177000
	fmt.Printf("Value: %v\n", statePopulation)
	fmt.Printf("Value: %v\n", statePopulation["Georgia"])
	fmt.Printf("\n")
	// take note that a map will not have a guaranteed order of elements
	// when appending data
	delete(statePopulation, "Georgia")
	fmt.Printf("Value: %v\n", statePopulation)
	fmt.Printf("Value: %v\n", statePopulation["Georgia"])
	fmt.Printf("\n")
	// notice how the element Georgia returns a 0 even when its deleted
	// This becomes a problem since it doesn't return an error
	// a way around this is by using the ok variable when assigning a value from a map
	geor, georExists := statePopulation["Georgia"]
	fmt.Printf("Georgia: %v, Exists: %v\n", geor, georExists)
	nev, nevExists := statePopulation["Nevada"]
	fmt.Printf("Georgia: %v, Exists: %v\n", nev, nevExists)
	fmt.Printf("\n")
	// when assigning a value to a var from a map, the value is a refference
	// just like a slice
	sp := statePopulation
	delete(sp, "Georgia")
	fmt.Printf("Value: %v\n", sp)
	fmt.Printf("Value: %v\n", statePopulation)
	fmt.Printf("\n")

	
	/*
	* STRUCTS
	*/
	type Hero struct {
		number					int						
		actorName				string
		companions			[]string
	}
	var drStrange = Hero {
		number: 20,
		actorName: "Bennedict Cumberbatch",
		companions: []string {
			"Spiderman",
			"Iron Man",
			"Thor",
		},
	}

	fmt.Printf("Value: %v\n", drStrange)
	fmt.Printf("Actor Name: %v\n", drStrange.actorName)
	fmt.Printf("Number: %v\n", drStrange.number)
	fmt.Printf("Companions: %v\n", drStrange.companions[1])
	fmt.Printf("\n")
	// unlike maps structs are independent structures
	// so when you assign a var from a different struct and make changes to it
	// its not going to affect the original struct
	a := drStrange
	a.actorName = "Tom Cruise"
	fmt.Printf("Original Actor Name: %v\n", drStrange.actorName)
	fmt.Printf("New Actor Name: %v\n", a.actorName)
	fmt.Printf("\n")

	// we also have embedding
	type Franchise struct{
		companyName					string
		heroes							[]Hero
	}
	marvel := Franchise {
		companyName: "Marvel",
		heroes: []Hero{
			drStrange,
		},
	}
	spiderman := Hero{
		actorName: "Tom Holland",
		number: 16,
		companions: []string{
			"Dr. Strange",
			"Iron Man",
			"Thor",
		},
	}
	marvel.heroes = append(marvel.heroes, spiderman)

	fmt.Printf("Marvel Franchise: %v\n", marvel)
	fmt.Printf("Marvel Heroes: %v\n", marvel.heroes)
	fmt.Printf("\n")
}