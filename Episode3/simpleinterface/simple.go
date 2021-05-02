package main

import "fmt"

type Articulo struct {
	Autor  string
	Titulo string
}

type Impresor interface {
	Strings() string
}

func Imprimir(i Impresor) {
	fmt.Println(i.Strings())
}

// Receivers in Go
func (a Articulo) Strings() string {
	return fmt.Sprintf("El autor %v escribió %v", a.Autor, a.Titulo)
}

func main() {
	a := Articulo{
		Autor:  "Alan, A. Donovan",
		Titulo: "The Go Programming Language",
	}
	//fmt.Println(a.Strings())

	Imprimir(a)
}
