package main

import (
	"fmt"
)

func OutputFunctions() {
	var i, j string = "Hello", "World"

	fmt.Print(i)
	fmt.Print(j)
	fmt.Print("\n")

	i, j = "Hello", "World"

	fmt.Println(i, j)

	i = "Hello"
	k := 15

	fmt.Printf("i has value: %v and type: %T\n", i, i)
	fmt.Printf("k has value: %v and type: %T", k, k)
	fmt.Print("\n")
}
