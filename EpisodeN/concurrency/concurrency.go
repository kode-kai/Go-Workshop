package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	// wg.Add(1)

	go func() {
		item := "order"
		process(item, c)
		order := <-c
		fmt.Printf("Processing %v %v\n", item, order)

		wg.Done()
	}()
	// go func() {
	// 	process("fund")
	// 	wg.Done()
	// }()

	wg.Wait()
}

func process(item string, c chan<- int) {
	for i := 1; i <= 5; i++ {
		c <- i
		time.Sleep(time.Millisecond * 500)
	}
}
