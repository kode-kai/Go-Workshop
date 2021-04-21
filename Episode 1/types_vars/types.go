package main

import "fmt"

// This is a constant variable
const maxFileSize = 1000000

var (
	Algo int
	Mas  string
)

func aliasDeclaration() {
	// Alias declaration
	type price float64
	var total price
	total = 899.99
	fmt.Println("Your total account is ->", total)
	fmt.Println("Conversion to Price ->", price(1527.55))

	// Getting data type and value
	fmt.Printf("Type: %T Value: %v\n", total, total)
	cadena := "diferencia"
	fmt.Printf("Esta cadena es: %v : %q : %T", cadena, cadena, cadena)
}

func operations() {
	// Same data type
	length, width := 10, 20
	fmt.Println("Area m2->", length*width)

	// Different data type (using type conversion or casting)
	height := 10.5
	fmt.Println("Area m3->", float64(length*width)*height)
}

func pointers() {
	// As a pointer
	myIntPtr := new(int)
	myFloatPtr := new(float32)
	myStringPtr := new(string)
	myBoolPtr := new(bool)

	*myIntPtr = 2013
	*myFloatPtr = 20.50
	*myStringPtr = "Hello there!"
	*myBoolPtr = true

	pointer := myIntPtr
	*pointer = 2014

	fmt.Println("My int var ->", *myIntPtr)
	fmt.Println("My float var ->", &myFloatPtr)
	fmt.Println("My string var ->", myStringPtr)
	fmt.Println("My bool var ->", *myBoolPtr)
	// Getting data type and value
	fmt.Printf("Type: %T Value: %v\n", maxFileSize, maxFileSize)
}

func types() {
	// Long way
	//var myInt int
	//var myFloat float32
	//var myString string
	//var myBool bool
	//
	//myInt = 2013
	//myFloat = 20.50
	//myString = "Hello there!"
	//myBool = true

	// Short way
	myInt := 2013
	myFloat := 20.50
	myString := "Hello there!"
	myBool := true

	fmt.Println("My int var ->", myInt)
	fmt.Println("My float var ->", myFloat)
	fmt.Println("My string var ->", myString)
	fmt.Println("My bool var ->", myBool)
}

func declarations() {
	// Declarations
	// long way
	var age int = 26
	// var age2 = 27
	var name string = "Adrian"
	fmt.Println("My age is:", age)
	fmt.Println("My name is:", name)

	// short way
	consoles := 3
	consoleName := "Xbox Series X"
	isTrue := true
	complex := 12 + 3i
	fmt.Println("I have:", consoles, "consoles")
	fmt.Println("One of them is:", consoleName)
	fmt.Println("One of them is:", isTrue)
	fmt.Println("One of them is:", complex)
}

func main() {
	//declarations()
	// types()
	pointers()
	// operations()
	// aliasDeclaration()
}
