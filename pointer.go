package main

import "fmt"


func Pointer(name string){
	var biodata string = name
	var poin *string = &biodata

	fmt.Println(biodata)
	fmt.Println(&biodata)
	fmt.Println(*poin)
	fmt.Println(&poin)
}