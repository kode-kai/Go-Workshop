package main

import (
	"fmt"
	"math/rand"
)

type RandomNumber struct {
	a int
	b int
}

var items int = 10

func main() {
	jobs := make(chan RandomNumber, items)
	results := make(chan int, items)

	// TODO watch the activity monitor
	go worker(jobs, results)
	// Create more goroutines to watch the performance

	for i := 0; i < items; i++ {
		random := RandomNumber{
			a: rand.Intn(100),
			b: rand.Intn(100),
		}
		fmt.Println(random.a, " ", random.b)
		jobs <- random
	}
	close(jobs)

	for i := 0; i < items; i++ {
		r := <-results
		fmt.Println("GCD is: ", r)
	}
}

func worker(jobs <-chan RandomNumber, results chan<- int) {
	for random := range jobs {
		results <- GCD(random.a, random.b)
	}
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
