package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	/*
	*		DEFER
	 */
	// * defer allows an execution to be executed right-
	// * before a function ends, before it returns
	// fmt.Printf("Defer function\n")
	// fmt.Printf("start, ")
	// defer fmt.Printf("middle, ")
	// fmt.Printf("last ")
	// fmt.Printf("\n")
	// * defer actually follows the last in first out principle
	// fmt.Printf("Defer function follows the last in first out principle:\n")
	// defer fmt.Printf("start, ")
	// defer fmt.Printf("middle, ")
	// defer fmt.Printf("last ")

	// fetchDataDefer()
	runServerPanic()
}

func fetchDataDefer() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", data)
}

func runServerPanic() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Println("Recieved Request")
		w.Write([]byte("Hello World!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}