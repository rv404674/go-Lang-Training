package main

import "fmt"
import "time"

func main(){

	i:=2

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	t:=time.Now()

	switch t{
	case t.Hour()<12:
		fmt.Println("it's before noon")
	default :
		fmt.Println("it's after noon")
	}
}