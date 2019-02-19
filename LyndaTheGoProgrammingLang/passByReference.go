package main

import "fmt"

func main(){
	a:=5
	changeA(a)
	fmt.Println("after pass as value:", a)

	changeAp(&a)
	fmt.Println("after passing pointer:", a)
}

func changeA(temp int){
	temp=10
}

func changeAp(temp *int){
	*temp = 10
}