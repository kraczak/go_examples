package main

import "fmt"

type person struct {
	// structs are mutable
	name    string
	surname string
	age     int
}

func newPerson(name string, surname string, age int) *person {
	p := person{name: name, surname: surname, age: age}
	return &p
}

func main() {
	fmt.Println(person{"Bob", "Johnson", 20})
	fmt.Println(person{name: "Bob", surname: "Johnson"})
	x := newPerson("Kornel", "Raczak", 25)
	x.age = 24
	fmt.Println(x)
}
