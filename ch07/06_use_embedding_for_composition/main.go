package main

import (
	"fmt"
)

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s, %s", e.Name, e.ID)
}

type Manager struct {
	Employee // an embedded field
	Reports  []Employee
}

func (m Manager) FindNewEmployees() []Employee {
	// do something
	return []Employee{}
}

type Inner struct {
	X int
}

type Outer struct {
	Inner
	X int // same name as Inner.X
}

func main() {
	m := Manager{
		Employee: Employee{
			Name: "Mustafa",
			ID:   "100bi@",
		},
		Reports: []Employee{},
	}

	fmt.Println(m.Employee.Name)
	fmt.Println(m.Name)
	fmt.Println(m.ID)
	fmt.Println(m.Description())

	o := Outer{
		Inner: Inner{
			X: 4,
		},
		X: 9,
	}
	fmt.Println(o.X)
	fmt.Println(o.Inner.X)

}
