package main

import "fmt"

func main() {

	s:= make([]string, 3)
	fmt.Println("emp", s)

	//We can set and get just with array
	s[0]="a"
	s[1]="b"
	s[2]="c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	fmt.Println("len:", len(s))

	s = append(s,"d")
	s = append(s, "d", "f")
	for i,j:=range s{
		fmt.Println(i)
		fmt.Println(j)
	}

	l := s[:3]
	fmt.Println(l)



}