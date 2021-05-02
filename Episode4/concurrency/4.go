package main

import "fmt"

func main() {
	c := make(chan string) //TODO: give it a buffer and send more messages

	// This operation will lock the execution
	// go getMessageConcurrently(c)
	c <- "Kode"

	message := <-c
	fmt.Println(message)
	// This will deadlock!
}

// func getMessageConcurrently(c chan string) {
// 	c <- "Kode"
// 	close(c)
// }
