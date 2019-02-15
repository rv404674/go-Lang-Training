package main

import "fmt"


type Salutation struct{
	name string
	greeting string
}

func CreateMessage(name string,  greeting ...string) (string, string){
	message = greeting[1] + " " + name
	alternate = "HEY!" + name

	return message, alternate
}

func Greet(salutation Salutation){
	message, alternate := CreateMessage(salutation.name, salutation.greeting, "yo")

	fmt.Println(message)
	fmt.Println(alternate)
}

func main(){
	var s = Salutation{"Bob", "hello"}
	Greet(s)
}