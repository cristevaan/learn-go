package main

import (
	"fmt"
	"strconv"
	"strings"
)

func cleanup() {
	fmt.Println("Cleanup...")
}

func configCheck(name string) {
	defer cleanup()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from error:", r)
		}
	}()
	if name == "" {
		panic("Name cannot be empty")
	}
	fmt.Println("Config name:", name)
}

func test() {
	fmt.Println("Test function")
}

func sayHello(name string) string {
	return "Hello " + name
}

func oddEven(number int) (string, int) {
	if number % 2 == 0 {
		return "Even", number
	} else {
		return "Odd", number
	}
}

func averageTwoNumber(a, b int) (result float64) {
	result = float64(a + b) / 2
	return
}

func addNumbers(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

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

	// map
	person := map[string]string{
		"name":    "Madilog",
		"address": "Indonesia",
	}
	fmt.Println(person)
	fmt.Println(person["name"])
	person["title"] = "Programmer"
	fmt.Println(person)
	delete(person, "address")
	fmt.Println(person)
	fmt.Println(len(person))

	var book = make(map[string]string)
	book["title"] = "Belajar Go"
	book["author"] = "Madilog"
	fmt.Println(book)

	// math operation
	a := 10
	fmt.Println(a + 5 - 2 * 3 / 2 % 2)
	a += 5
	fmt.Println(a)
	a -= 2
	fmt.Println(a)
	a *= 3
	fmt.Println(a)
	a /= 2
	fmt.Println(a)
	a %= 2
	fmt.Println(a)
	a++
	fmt.Println(a)
	a--
	fmt.Println(a)

	// comparison & logical operation
	fmt.Println(a == 10)
	fmt.Println(a != 10)
	fmt.Println(a > 10)
	fmt.Println(a >= 10)
	fmt.Println(a < 10)
	fmt.Println(a <= 10)

	fmt.Println(a > 5 && a < 15)
	fmt.Println(a > 5 || a < 5)
	fmt.Println(!(a > 5 && a < 15))

	// if statement
	if a > 10 {
		fmt.Println("a lebih besar dari 10")
	} else if a == 10 {
		fmt.Println("a sama dengan 10")
	} else {
		fmt.Println("a lebih kecil dari 10")
	}

	if b := a * 2; b > 15 {
		fmt.Println("b lebih besar dari 15")
	} else {
		fmt.Println("b tidak lebih besar dari 15")
	}

	// switch case
	switch a {
	case 8:
		fmt.Println("a adalah 8")
	case 9:
		fmt.Println("a adalah 9")
	case 10:
		fmt.Println("a adalah 10")
	default:
		fmt.Println("a bukan 8, 9, atau 10")
	}

	switch c := a * 2; {
	case c < 15:
		fmt.Println("c lebih kecil dari 15")
	case c == 20:
		fmt.Println("c sama dengan 20")
	default:
		fmt.Println("c lebih besar dari 15 dan tidak sama dengan 20")
	}

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println("Perulangan ke-", i)
	}

	j := 5
	for j < 9 {
		fmt.Println("Perulangan ke-", j)
		j++
	}

	food := []string{"Nasi Goreng", "Mie Goreng", "Sate"}
	for index, food := range food {
		fmt.Printf("Makanan %d: %s\n", index, food)
	}

	for k := 0; k < 6; k++ {
		if k == 1 {
			continue
		}
		if k == 5 {
			break
		}
		fmt.Println("Perulangan ke-", k)
	}

	// function
	test()

	fmt.Println(sayHello("Madilog"))

	hasil, number := oddEven(10)
	fmt.Printf("%d adalah bilangan %s\n", number, hasil)

	avg := averageTwoNumber(10, 20)
	fmt.Printf("Rata-rata: %.2f\n", avg)

	total := addNumbers(10, 20, 30, 40, 50)
	fmt.Printf("Total: %d\n", total)

	// defer, panic, & recover
	configCheck("MyApp")
	fmt.Println("Program selesai")
	configCheck("")
	fmt.Println("Program selesai")
}