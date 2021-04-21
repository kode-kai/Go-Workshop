package main

import (
	"fmt"
	"strings"
)

// we receive all the variadic numbers with the ... operator before the type
func sum(numbers ...int) int {
	var total int
	for _, v := range numbers {
		total += v
	}
	return total
}

func printer(fruit string, fruit2 string, fruits ...string) string {
	var separator string = " "
	fruits = append(fruits, fruit2)
	fruit += separator + strings.Join(fruits, separator)
	return fruit
}

func variadicAnyType(any ...interface{}) {
	for _, v := range any {
		fmt.Printf("%v of type %T\n", v, v)
	}
}

func main() {
	fmt.Println(sum(1, 2, 3, 4))
	numbers := []int{1, 2, 3, 5}
	// In case we do have a slice we need to use ... operator after the slice
	fmt.Println(sum(numbers...))

	fmt.Println(printer("apple", "watermellon", "orange", "grapes"))

	/// Variadic function of any type
	variadicAnyType(1, "red", true, 10.5, []string{"foo", "bar", "baz"},
		map[string]int{"adrian": 26, "antonio": 30})
}
