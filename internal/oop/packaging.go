package oop

import "fmt"

type Animal struct {
	Name string
}

func (a *Animal) Eat() {
	fmt.Printf("%s eat...\n", a.Name)
}

func (a *Animal) Hi() {
	fmt.Printf("%s say hi\n", a.Name)
}
