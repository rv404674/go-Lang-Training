package main

import "fmt"

func main(){
	var message string = "hello world"

	var greeting *string = &message
	*greeting = "rahul verma"
	fmt.Println(message, *greeting)

}