package main

import (
	"fmt"
	"strconv"
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

	// type conversion
	value1 := 32768.765
	value2 := int(value1)

	fmt.Println(value1)
	fmt.Println(value2)

	value3 := 100
	value4 := strconv.Itoa(value3)

	fmt.Printf("value3 is of type %T\n", value3)
	fmt.Printf("value4 is of type %T\n", value4)

	value5 := "100"
	value6, _ := strconv.Atoi(value5)

	fmt.Printf("value5 is of type %T\n", value5)
	fmt.Printf("value6 is of type %T\n", value6)

	value7 := true
	value8 := strconv.FormatBool(value7)
	value9, _ := strconv.ParseBool(value8)

	fmt.Printf("value7 is of type %T\n", value7)
	fmt.Printf("value8 is of type %T\n", value8)
	fmt.Printf("value9 is of type %T\n", value9)

}