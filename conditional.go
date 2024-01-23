package main

import "fmt"

func Conditional(username string) {
	if username == "gabriel" {
		fmt.Println("Anda Berhasil Login!")
	} else {
		fmt.Println("Username or Password wrong!")
	}
}