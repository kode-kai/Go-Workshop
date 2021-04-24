package main

import (
	"fmt"
	"sort"
)

func sortPrimitives() {
	// Sorting a slice of strings
	s := []string{"find", "a", "way", "to", "sort"}
	sort.Strings(s)
	fmt.Println(s)

	// Sorting a slice of ints
	i := []int{3, 12, 66, 5, 1}
	sort.Ints(i)
	fmt.Println(i)
}

func sortStruct() {
	type Family struct {
		name string
		age  int
	}

	family := []Family{
		{"Alice", 23},
		{"David", 2},
		{"Eve", 2},
		{"Bob", 25},
	}

	sort.SliceStable(family, func(i, j int) bool {
		return family[i].age < family[j].age
	})
	fmt.Println(family)
}

func sortMap() {
	m := map[string]int{"Alice": 2, "Cecil": 1, "Bob": 3}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, m[k])
	}
}

func main() {
	sortPrimitives()
	sortStruct()
	sortMap()
}
