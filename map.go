package main

import "fmt"

func Map(){
	biodata := map[string]string {
		"firstname": "Gabriel",
		"lastname": "Putra",
		"pendidikan": "UNAI",
		"umur": "24",
	}

	fmt.Println("My Name:",biodata["firstname"],biodata["lastname"])
	fmt.Println(biodata["pendidikan"])
}