package main

import "fmt"

type Employee struct {
	name                 string
	age                  int
	lastName             string
	programmingLanguages []string
	isSenior             bool
}

type Employees []Employee

func declarations() {
	// Long Way
	var employeeSenior Employee
	employeeSenior = Employee{
		name:                 "Antonio",
		lastName:             "Feregrino",
		age:                  39,
		programmingLanguages: []string{"C#", "Python"},
		isSenior:             true,
	}
	fmt.Println(employeeSenior)

	// Short way
	employee := Employee{
		name:                 "Adrian",
		lastName:             "Fernandez",
		age:                  26,
		programmingLanguages: []string{"Go", "JavaScript", "TypeScript", "C#"},
		isSenior:             false,
	}
	fmt.Println(employee)

	// Changing a property from the struct
	employee.age = 30
	employee.isSenior = true

	for _, language := range employee.programmingLanguages {
		if language == "Go" {
			fmt.Println(language, " is my favorite language")
		}
		fmt.Println(language)
	}

}

func pointers() *Employees {
	employee := new(Employee)
	*employee = Employee{
		name:                 "Adrian",
		lastName:             "Fernandez",
		age:                  26,
		programmingLanguages: []string{"Go", "JavaScript", "TypeScript", "C#"},
		isSenior:             false,
	}
	employeeSenior := Employee{
		name:                 "Antonio",
		lastName:             "Feregrino",
		age:                  39,
		programmingLanguages: []string{"C#", "Python"},
		isSenior:             true,
	}
	fmt.Println(*employee)

	employees := Employees{}
	employees = append(employees, *employee)
	employees = append(employees, employeeSenior)
	return &employees
}

func structAsReference() *Employee {
	e := Employee{}
	return &e
}

func assignNameAsReference(e *Employee) {
	if e.name == "" {
		e.name = "Ad"
	}
}

func main() {
	// employees := pointers()
	// fmt.Println(employees)

	e := structAsReference()
	assignNameAsReference(e)
	fmt.Println(e.name)
}
