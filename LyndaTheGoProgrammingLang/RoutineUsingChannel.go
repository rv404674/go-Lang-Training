package main

import (
	"fmt"
)

func printName(x string, y int){
	i:=0
	for i<y{
		fmt.Println(x)
		i++
	}
}

func main(){

	// When this go routine is done we would like it to send a message to main thread and say that i am done. We are gonna wait until
	// we get that message

	done := make(chan bool)

	//anoymous fn
	go func(){
		//into this block sequential exec will be performed
		printName("rahul",3)
		done <- true
	}()

	printName("sachin",3)
	<- done // it will basically wait until it reads from done.

}