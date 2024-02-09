package main

import "fmt"


func Map(){
	biodata := map[string]string{
		"firstname": "Gabriel",
		"lastname": "Putra",
		"pendidikan": "UNAI",
		"umur": "24",
	}
	data := map[string]map[string]string{
		"Content-Disposition": {
			"name": "picture",
			"filename": "algorism-logo.png",
		},
		"Content-Type": {
			"0":"image/png",
		},
	}
	fmt.Println("My Name:",biodata["firstname"],biodata["lastname"])
	fmt.Println(biodata["pendidikan"])
	fmt.Println(biodata)
	fmt.Println(data["Content-Type"]["0"])
}