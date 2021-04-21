package main

import (
	"fmt"
	"strconv"
)

var Accesible string = "2"

func main() {
	numberString := "1000"
	number, err := strconv.Atoi(numberString)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Hello Kode Kai!", number)
}
