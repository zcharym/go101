package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

// return a pointer to local variable as local variable
func newPerson(name string, age int) *person {
	p := person{age: age, name: name}
	return &p
}

func main() {
	fmt.Println(newPerson("Jon", 23))

	p := newPerson("Tom", 22)
	fmt.Println(p.age)
	p.age = 10
	fmt.Println(p.age)
	sp := &p
	fmt.Println(sp)

}
