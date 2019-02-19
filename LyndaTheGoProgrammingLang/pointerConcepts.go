package main

import "fmt"

func main(){

	//Declare variable of type int
	count := 10

	//Display the "value of" and "address"
	fmt.Println("Before:", count, &count)

	//Pass the "address of" the variable 
	increment(&count)

	fmt.Println("after: ", count, &count)
}

//increment declares count as a pointer
func increment(inc *int){

	// Increment the value that the "pointers"
	// (de-referencing)
	*inc++
	fmt.Println("Inc: ",*inc, &inc, inc)
}