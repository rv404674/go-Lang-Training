package main

import "fmt"

func main(){
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	var c chan int
	fmt.Printf("Chan: %v\n", c)

	var m map[int]int
	fmt.Printf("Map: %v\n", m)

	var s1 []int
	fmt.Printf("Slice: %v\n", s1)

}