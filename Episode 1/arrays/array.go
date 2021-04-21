package main

import "fmt"

func declarations() {
	// Array declarations
	// Long way
	var myArray [5]int
	myArray[0] = 2
	myArray[1] = 4
	myArray[2] = 6
	myArray[3] = 8
	myArray[4] = 10
	// Short way
	myArray2 := [10]int{}
	myArray2[5] = 1
	// Short way with composite literal
	myArray3 := [15]int{2, 3, 9, 12, 15}
	myArray3[8] = 1000

	myArrayString := [100]string{"cadena1", "cadena2", "cadena3"}

	fmt.Println("Array 1 ->", myArray)
	fmt.Println("Array 2 ->", myArray2)
	fmt.Println("Array 3 ->", myArray3)
	fmt.Println("Array 4 ->", myArrayString)

}

func pointers() {
	// As a pointer
	myArray := new([20]int)
	fmt.Println("Array 4 ->", myArray)
}

func iterating() {
	myArray := [5]int{1, 2, 3, 4, 5}
	// Iterating Array
	for i, v := range myArray {
		fmt.Println("The array in the position", i, "contains", v)
	}
}

func main() {
	declarations()
	//pointers()
	//iterating()
}
