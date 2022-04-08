package main

import (
	"fmt"
	"strconv"
)

func main() {

	/*
		* BOOLEANS *
	*/
	fmt.Printf("\nBOOLEANS\n")
	isYoung := true
	fmt.Printf("is young: %v %T", isYoung, isYoung)
	
	
	var defaultBool bool
	fmt.Printf("\nDefault Boolean value: %v %T", defaultBool, defaultBool)
	
	n := 1 == 1
	m := 2 == 1
	
	fmt.Printf("\n1 is equal to 1: %v \n2 is equal to 1: %v", n, m)
	
	/*
	* INTEGERS *
	*/
	fmt.Printf("\nINTEGERS\n")
	var defaultInt int
	fmt.Printf("\nDefault value: %v %T", defaultInt, defaultInt)
	// int 8 from -128 to 127
	// int 16 from -32,768 to 32,767
	// int 32 from -2,147,483,648 to -2,147,483,647
	var unsignedInt uint = 12
	fmt.Printf("\nValue: %v %T\n", unsignedInt, unsignedInt)
	// uint 8 from 0 to 255
	// uint 16 from 0 to 65,535
	// uint 32 from 0 to 4,294,967,295
	
	// cannot add different int types
	// var a int = 3
	// var b int8 = 10
	// !fmt.Println(a+b)
	
	var a int = 10
	var b int = 3
	fmt.Println(float32(a) / float32(b))
	
	
	// bit operators - this checks which has its bit set using the comparison operator
	// in c & d it only checks which ones have a bit set which would result to 0010
	// 0010 is 2 in binary
	c := 10 // 1010
	d := 3	// 0011
	fmt.Println(c & d) // 0010
	fmt.Println(c | d) // 1011
	fmt.Println(c ^ d) // 1001
	fmt.Println(c &^ d) // 0100

	// bit shifting
	e := 7 // 2^3
	fmt.Println(e << 3) // 2^3 * 2^3 = 2^6
	fmt.Println(e >> 3) // 2^3 / 2^3 = 2^0


	/*
	* STRINGS *
	*/
	fmt.Printf("\nSTRINGS\n")
	s := "abcdefghijklmnopqrstuvwxyz "
	fmt.Printf("s[0]: %v, type of s[0]:%T, string of s[0]:%v, type of s[0]:%T\n", s[0], s[0], string(s[0]), s[0])
	fmt.Printf("%v\n", s)
	// strings are immutable
	// s[0] = "T"
	// will not work
	
	// you can also convert strings to bytes
	bts := []byte(s)
	fmt.Printf("%v, %T\n", bts, bts)

	/*
	* RUNES *
	*/
	fmt.Printf("\nRUNES\n")
	// make sure to read the documentation first before using runes
	// runes are single characters
	r := 'a'
	fmt.Printf("%v, %T\n", r, r)

	/*
	* VARIABLES
	*/
	var (
		firstName = "cloyd"
		lastName = "abad"
		age1 = 20
	)

	// Upper cased variables are accessible outside of the package or file
	// var I = 20
	// Lower cased variables are only accessible inside the file
	
	fmt.Println(firstName, lastName, age1)
	
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