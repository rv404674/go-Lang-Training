package main

import "fmt"

type Salutation struct{
	name string
	greeting string
}

func main(){


	var s []int
	s= make([]int, 3)
	//s = make([]int, 3, 10)
	//(, initial length, capacity)

	// or s:= int[] {1,10,500,25}

	s[0] = 1
	s[1] = 10
	s[2] = 50

	slice := []Salutation{
		{"Bob","Hello"},
		{"joe", "Hi"},
		{"Mary", "What is up"},
	}

	slice = slice[0:2] // 1 but not 2
	// slice = slice[1:len(slice)]
	// slice = slice[1:] (both will work as same)


	for i,j := range slice{
		fmt.Println(i)
		fmt.Println(j.name + " " + j.greeting)
	}

	//APPENDING
	slice = append(slice, Salutation{"Frank","Hi"})

	for i,j := range slice{
		fmt.Println(i)
		fmt.Println(j.name + " " + j.greeting)
	}


}