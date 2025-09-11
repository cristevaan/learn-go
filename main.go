package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println("Hello World!")

	// tipe data
	var name string = "Madilog"
	age := 80
	const isAlive = true

	fmt.Printf("Name: %v\n", name)
	fmt.Printf("Age: %v\n", age)
	fmt.Printf("Is Alive: %v\n", isAlive)

	text := `Ada dua kemungkinan
	1. Kemungkinan
	2. Kemungkinan`

	fmt.Println(text)

	fmt.Println(len(name))
	fmt.Println("Hello " + strings.ToUpper(name))
}