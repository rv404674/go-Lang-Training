package main

import "fmt"

func main(){

	// Here is a basic example
	if 7%2 ==0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is even")
	}

	// a statement can preced with a; any variable declared in this statement wil be available in other as well
	if a:=9; a<0 {
		fmt.Println(a," is negative")
	} else if a < 10 {
		fmt.Println(a, " has 1 digit")
	} else {
		fmt.Println(a, " has multiple digits")
	}


}