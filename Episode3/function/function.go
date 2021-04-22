package main

import (
	"fmt"
	"strconv"
)

type fn func(string) int

func function() {
	fmt.Println("This is just a function!")
	return
}

func receiveReturn(a int) int {
	return a
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

func receiveFunction(f fn) fn {
	number := f("Try me!")
	fmt.Println(number)
	return f
}

func main() {
	// function()
	// receiveReturn(12)
	// returnMultiple()

	f := func(s string) int { return 5 }
	receiveFunction(f)

	func(s string) {
		fmt.Println(s)
	}("I'm an anonymous function!")
}
