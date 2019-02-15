package main

import "fmt"


type Salutation struct{
	name string
	greeting string
}



func main(){

	slice:= [] Salutation{
		{"Bob", "hello"},
		{"Joe", "Hi"},
		{"Mary", "gay"},
	}

	for i,j := range slice{
		fmt.Println(i)
		fmt.Println(j.name + " " + j.greeting)
	}

}