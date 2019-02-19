package main

import "fmt"

type Person struct{
	age int
}

func (person Person) SetAge (temp int){
	person.age = temp
}

func (person Person) IsMajor() bool{
	return person.age > 18
}

func main(){
	p := Person{15}
	fmt.Println(p.IsMajor())
}