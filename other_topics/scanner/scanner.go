package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	element := input.Text()
	number, _ := strconv.Atoi(element)

	fmt.Println(element)
	fmt.Println("int: ", number)
}
