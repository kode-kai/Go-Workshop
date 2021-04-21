package main

import "fmt"

func main() {
	fmt.Println("Conexion a BD")

	defer func() {
		message := recover()
		fmt.Println(message)
	}()

	panic("Error it panics")
}
