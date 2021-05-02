package main

import (
	"fmt"
	"math/rand"
)

type RandomNumber struct {
	a int
	b int
}

func main() {
	jobs := make(chan RandomNumber, 10)
	results := make(chan int, 10)

	go worker(jobs, results)

	for i := 0; i < 10; i++ {
		random := RandomNumber{
			a: rand.Intn(100),
			b: rand.Intn(100),
		}
		fmt.Println(random.a, " ", random.b)
		jobs <- random
	}
	close(jobs)

	for i := 0; i < 10; i++ {
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
