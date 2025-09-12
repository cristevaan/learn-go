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

	// array
	var scores [3]int = [3]int{90, 80, 70}
	fmt.Println(scores)
	fmt.Println(scores[0])
	scores[0] = 100
	fmt.Println(scores[0])
	fmt.Println(scores)
	fmt.Println(len(scores))

	names := [...]string{"Madilog", "Budi", "Joko"}
	for i := 0; i < len(names); i++ {
		fmt.Printf("Index %d: %s\n", i, names[i])
	}

	for i, name := range names {
		fmt.Printf("Index %d: %s\n", i, name)
	}

	for _, name := range names {
		fmt.Printf("Name: %s\n", name)
	}

	arrx := [5]int{1: 10, 3: 30}
	fmt.Println(arrx)

	// slice
	var fruits = []string{"Apple", "Banana", "Cherry"}
	fmt.Println(fruits)
	fmt.Println(len(fruits))
	fmt.Println(cap(fruits))
	fruits = append(fruits, "Durian")
	fmt.Println(fruits)
	fmt.Println(len(fruits))
	fmt.Println(cap(fruits))

	myarr := [5]int{10, 20, 30, 40, 50}
	myslice := myarr[1:4]
	fmt.Println(myslice)
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))

	myslice1 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice2 := make([]int, 3)
	copy(myslice2, myarr[1:4])
	fmt.Printf("myslice2 = %v\n", myslice2)
	fmt.Printf("length = %d\n", len(myslice2))
	fmt.Printf("capacity = %d\n", cap(myslice2))

}