package main

import (
	"fmt"
	"time"
)

func main() {
	// TODO 1: make orders concurrent
	process("orders")

	// TODO 2: make funds concurrent
	process("funds")
}

func process(item string) {
	for i := 1; ; i++ {
		fmt.Printf("Processing %v %v\n", i, item)
		time.Sleep(time.Millisecond * 500)
	}
}
