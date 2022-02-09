package oop

import "fmt"

type Cat struct {
	*Animal
	Color string
}

func (c *Cat) Show() {
	fmt.Printf("%s color is %s\n", "a", c.Color)
}
