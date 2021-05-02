package main

import (
	"fmt"
	"time"
)

func main() {
	var item string = "orders"
	c := make(chan int)

	go process(item, c)
	for {
		// TODO 1: Get ok in order to know when finish the loop
		orderId := <-c
		fmt.Printf("Processing %v %v\n", orderId, item)
	} // TODO 2: Iterate over the range channel
}

func process(item string, c chan int) {
	for i := 1; i <= 5; i++ {
		c <- i
		time.Sleep(time.Millisecond * 500)
	}
}
