package main

import "fmt"

func declarations() {
	// Slice declarations
	// Long way
	var mySlice []int
	//mySlice[0]=2 // THIS WILL PANIC (The zero value of a slice is nil)
	// Short way
	mySlice2 := []int{}
	// Short way with composite literal
	mySlice3 := []int{2, 3, 9, 12, 15}

	// As a pointer
	mySlice4 := new([]int)
	*mySlice4 = make([]int, 5)
	(*mySlice4)[2] = 2
	// Using Make
	// mySlice5 := make([]int, 5)
	mySlice5 := make([]int, 10)
	fmt.Println("Slice 1 ->", mySlice)
	fmt.Println("Slice 2 ->", mySlice2)
	fmt.Println("Slice 3 ->", mySlice3)
	fmt.Println("Slice 4 ->", mySlice4)
	fmt.Println("Slice 5 ->", mySlice5)
}

func myAppend() {
	var mySlice []int
	mySlice = make([]int, 2)

	fmt.Println(len(mySlice), cap(mySlice))
	mySlice[0] = 1
	mySlice[1] = 2
	mySlice = append(mySlice, 3)
	mySlice = append(mySlice, 4)
	mySlice = append(mySlice, 5)
	// fmt.Println(cap(mySlice))
	mySlice = append(mySlice, 6)
	fmt.Println(len(mySlice), cap(mySlice))
	// fmt.Println(cap(mySlice))

	fmt.Println(mySlice)
}

func copy() {
	cities := []string{"New York", "London", "Chicago", "Beijing", "Delhi", "Mumbai", "Bangalore", "Hyderabad", "Hong Kong"}
	newCities := cities

	citiesReference := new([]string)
	*citiesReference = make([]string, 100)
	citiesReference = &cities

	cities = append(cities, "Mexico")
	fmt.Println(cities)
	fmt.Println(newCities)
	fmt.Println(*citiesReference)
	*citiesReference = append(*citiesReference, "Londres")
	fmt.Println(cities)
}

func insert(num int, pos int, numbers []int) []int {
	var result []int
	for i, _ := range numbers {
		if i == pos {
			result = append(result, numbers[:i]...)
			result = append(result, num)
			result = append(result, numbers[i:]...)
			break
		}
	}
	return result
}

func main() {
	// declarations()
	// myAppend()
	// numbers := []int{1, 2, 3}
	// result := insert(3, 1, numbers)
	// fmt.Println(result)
	// copy()
}
