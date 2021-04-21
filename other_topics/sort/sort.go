package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people)

	// The other way is to use sort.Slice with a custom Less
	// function, which can be provided as a closure. In this
	// case no methods are needed. (And if they exist, they
	// are ignored.) Here we re-sort in reverse order: compare
	// the closure with ByAge.Less.
	sort.Slice(people, func(a, b int) bool {
		return people[a].Name < people[b].Name
	})
	fmt.Println(people)

	numbers := []int{1, -3, 100, 33, -12, 4}
	sort.SliceStable(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
	fmt.Println(numbers)

}
