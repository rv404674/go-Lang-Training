package main

import "fmt"


//Capital "S" of salutation means this is publicy visisble inside this package, lower case is not visible
type Salutation struct{
	name string
	greeting string
}

// anything inside this is constant

const {
	PI = 3.14
	Language = "GO"
	//iota represents successive untyped int constants
	A = iota
	B = iota
	C = iota
}

func main(){
	var s =  Salutation{"Rahul Verma", "Hello world!"}
	// Salutation{greeting:"Hello!", name:"Rahul Verma"}
	//
	// var s = Salutation{}
	// s.name = "Bob" 
	// s.greeting = "Hello World"

	fmt.Println(s.name)
	fmt.Println(a, b, c)
	fmt.Println(s.greeting)
}