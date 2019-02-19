package main

import "fmt"

type Speaker interface{
	speak()
}

type Dog struct{}

func (d Dog) speak() {
	fmt.Println("bow bow")
}

type Car struct{}

func (c Car) speak() {
	fmt.Println("honk honk")
}

type Human struct{}

func (h Human) speak() {
	fmt.Println("hello")
}

func main(){
	speakers:= [] Speaker{
		Dog{},
		Car{},
		Human{},
	}

	for _,sp := range speakers{
		sp.speak()
	}
}