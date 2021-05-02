package main

import (
	"fmt"
	"time"
)

func main() {
	// var wg sync.WaitGroup
	// TODO: wait until finish both process calls
	go process("orders")

	go process("funds")

	// Workarounds examples
	// fmt.Scanln()
	// time.Sleep(time.Second * 2)
}

func process(item string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Processing %v %v\n", i, item)
		time.Sleep(time.Millisecond * 500)
	}
}
