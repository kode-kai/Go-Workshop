package main

import "fmt"

func ifElse() {
	number := 23
	if number%2 == 0 {
		fmt.Print(number, " even number")
	} else {
		fmt.Println(number, " odd number")
	}

	m := make(map[string]int)
	m["test"] = 100
	// declaration within if statement
	if value := m["test"]; value > 99 {
		fmt.Println("Es mayor a 99!")
	}
}

func switchs() {
	number := 100
	// BY VALUE
	switch number {
	case 1:
		fmt.Println(number, " is equal to 1")
	case 100:
		fmt.Println(number, " is equal to 100")
	}

	// BY CONDITION
	switch {
	case number > 80:
		fmt.Println(number, " is greater than 80")
	case number < 20:
		fmt.Println(number, " is greater than 20")
	default:
		fmt.Println(number, " didn't match any special case")
	}

	// Fallthrough
	switch 1 {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
		fallthrough
	case 3:
		fmt.Println("3")
	}

	// Label with Exit Break
Loop:
	for _, ch := range "a bc\n" {
		switch ch {
		case ' ': // skip space
			break
		case '\n': // break at newline
			break Loop
		default:
			fmt.Printf("%c\n", ch)
		}
	}
}

func forLoops() {
	// FOR - CLASSIC
	sum := 0
	for i := 1; i < 6; i++ {
		sum += i
	}
	fmt.Println("Sum ->", sum) // 10 (1+2+3+4)

	// FOR - AS WHILE
	n := 1
	for n < 5 {
		n *= 2
	}
	fmt.Println("N ->", n) // 8 (1*2*2*2)

	// FOR - INFINITE
	// for {
	// 	fmt.Println("Esto es infinito")
	// 	continue
	// }

	// FOR - AS A FOR EACH
	strings := []string{"hello", "world"}
	for i, v := range strings {
		fmt.Println(i, v)
	}
	for i, ch := range "HELLO" {
		fmt.Printf("%#U starts at byte position %d\n", ch, i)
	}

	// FOR - EXIT LOOP
	sum = 0
	for i := 1; i < 5; i++ {
		if i%2 != 0 { // skip odd numbers
			continue
		}
		sum += i
	}
}

func main() {
	// ifElse()
	switchs()
	// forLoops()
}
