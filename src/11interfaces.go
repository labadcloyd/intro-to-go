package main

import (
	"fmt"
	"math"
)

func main() {
	///* INTERFACES
	// with interfaces you can group similar methods
	// for example, with the circle and rectangle structs
	// they both have a method of area which we can group
	// through an interfaces and use it however we like
	c1 := circle{radius: 2.1}
	r1 := rect{width: 5.1, height: 1.1}
	shapes := []shape{c1, r1}
	fmt.Println(shapes[0].area(),shapes[1].area())

	for _, v := range shapes {
		fmt.Println(v.area())
	}
}

type shape interface {
	area()	float64
}
type circle struct {
	radius float64
}
type rect struct {
	width  float64
	height float64
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (r rect) area() float64 {
	return r.width * r.height
}
