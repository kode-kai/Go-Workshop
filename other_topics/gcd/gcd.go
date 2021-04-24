package main

import "fmt"

func GCD(a, b int) int {
	fmt.Printf("a: %v b: %v\n", a, b)
	for b != 0 {
		a, b = b, a%b
		fmt.Printf("a: %v b: %v\n", a, b)
	}
	return a
}

func main() {
	fmt.Println(GCD(20, 5))
}
