package main

import "fmt"

func main(){
	prefixMap := map[string]string{
		"Bob": "mr",
		"Joe":"Dr",
		"mary":"Mrs",
	}

	name := "mary"
	delete(prefixMap,"mary")
	//prefixMap["mary"] = ""

	for i,j := range prefixMap{
		fmt.Println(i,j)
	}

	if value, exists := prefixMap[name]; exists{
		fmt.Println("exists",value)
	} else{
		fmt.Println("doesnt exists",value)
	}

}