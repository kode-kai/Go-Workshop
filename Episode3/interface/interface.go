package main

import (
	"fmt"
)

type Calculator struct {
	feature NewFeature
}

func New(f NewFeature) *Calculator {
	return &Calculator{f}
}

type NewCalculator interface {
	Add(int, int) int
	Substract(int, int) int
	Multiply(int, int) int
	Divide(int, int) int
	Sqrt(int) float64
	Exp(int, int) int64
}

func anyType(a ...interface{}) {
	for i, v := range a {
		fmt.Println(i, ": ", v)
	}
}

func main() {
	// numbers := []int{1, 3, 5, 7}
	// anyType(numbers)

	f := NewF()
	c := New(f)
	sum := c.Add(3, 8)
	sqrt := c.Sqrt(25)
	fmt.Println(sum)
	fmt.Println(sqrt)
}

func (c *Calculator) Add(a int, b int) int {
	return a + b
}

func (c *Calculator) Substract(a int, b int) int {
	return a - b
}

func (c *Calculator) Multiply(a int, b int) int {
	return a * b
}

func (c *Calculator) Divide(a int, b int) int {
	return a / b
}

func (c *Calculator) Sqrt(n int) float64 {
	return c.feature.SqrtF(n)
}

func (c *Calculator) Exp(x int, n int) int64 {
	return c.feature.ExpF(x, n)
}
