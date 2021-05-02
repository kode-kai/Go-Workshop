package main

import (
	"fmt"
	"strconv"
)

func function() {
	fmt.Println("This is just a function!")
}

func receiveReturn(a int) int {
	return a
}

func receiveReturn2(a, b int, isOn, isOff bool) int {
	return a + b
}

func returnMultiple() (int, bool, string) {
	isSomething := true
	s := "34"

	value, err := strconv.Atoi(s)
	if err != nil {
		return 0, false, ""
	}

	return value, isSomething, s
}

type fn func(string) int

func receiveFunction(f fn) fn {
	number := f("213")
	fmt.Println(number)
	return f
}

func main() {
	// function()

	// receiveReturn(12)
	// receiveReturn2(12, 24)
	// returnMultiple()

	f := func(s string) int {
		number, err := strconv.Atoi(s)

		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("Hubo un error en la conversion")
		}

		return number
	}
	f1 := receiveFunction(f)
	f1("Try me again")
	// fmt.Println(n)
	// s := "I'm an anonymous function!"
	// func(s string) {
	// 	fmt.Println(s)
	// }(s)
}
