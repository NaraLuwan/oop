package oop

import "fmt"

// 封装特性

type Animal struct {
	Name string
}

func (a *Animal) eat() {
	fmt.Printf("%s eat...\n", a.Name)
}
