package main

import "fmt"

func goodnight(s string){
	fmt.Println("goodnight")
}
func main(){

	var s = "good night"
	defer goodnight(s)
	fmt.Println("Good morning")
}