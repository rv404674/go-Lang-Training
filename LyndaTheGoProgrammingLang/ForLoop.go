package main

import "fmt"

func main(){

	//There is only one type of loop in go i.e for

	a := 5
	for i:=0; i<a ; i++{
		fmt.Println("hello world")
	}

	//use for loop as a while loop

	i:=0
	for i<a {
		fmt.Println("hello")
		i++
	}

	//infinite loop
	for {
		break

	}

	i=0
	for {
		fmt.Println("hello")
		i++
		if i==2{
			break
		}
	}
}