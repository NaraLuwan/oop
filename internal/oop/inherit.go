package oop

import "fmt"

// 继承特性

type Cat struct {
	Animal
	color string
	age   int
}

func (c *Cat) show() {
	fmt.Printf("%s color is %s\n", "a", c.color)
}
