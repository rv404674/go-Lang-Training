package main

import "fmt"


type Salutation struct{
	name string
	greeting string
}

func CreateMessage(name, greeting string) (string, string){
	return greeting + " " + name, "HEY!" + name
}

func Greet(salutation Salutation){
	_, alternate := CreateMessage(salutation.name, salutation.greeting)

	fmt.Println(alternate)
}

func main(){
	var s = Salutation{"Bob", "hello"}
	Greet(s)
}