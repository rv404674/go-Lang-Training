package main

import (
	"fmt"
	"time"
)

func printName(x string, y int){
	i:=0
	for i<y{
		fmt.Println(x)
		i++
	}
}

func main(){

	// this function will be running in its own thread
	go printName("rahul",3) // this linne is not blocking as outcome of this thread in coming after sachin 
	printName("sachin",3)

	time.Sleep(100 * time.Millisecond)
}