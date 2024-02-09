package main

import (
	"fmt"
	"strings"
)

func Strings() {
	arr := []string{"a", "b", "c", "d", "e"}

	// Mengecek apakah "c" terdapat dalam arr
	terdapat := strings.Contains(strings.Join(arr, ","), "k")

	fmt.Println("Apakah \"c\" terdapat dalam arr?", terdapat) // true
	fmt.Println(strings.Join(arr, ","))
	fmt.Println(arr)
}