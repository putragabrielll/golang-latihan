package main

import "fmt"


func Looping(number int){
	for i := 1; i < number + 1; i++ {
		fmt.Println("Looping ke- ", i)
	}
}