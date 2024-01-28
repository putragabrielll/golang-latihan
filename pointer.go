package main

import "fmt"


func Pointer(name string){
	var biodata string = name
	var poin *string = &biodata

	fmt.Println(biodata)
	fmt.Println(&biodata) // "&" untuk cek memori keberapa dia terpakai 
	fmt.Println(*poin) // "*" untuk cek isi dari suatu variable
	fmt.Println(&poin)
}