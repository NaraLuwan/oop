package main

import (
	"github/NaraLuwan/oop/internal/oop"
)

func main() {
	// 封装
	peppa := &oop.Animal{Name: "peppa"}
	peppa.Eat()

	// 继承
	tom := &oop.Cat{
		Animal: &oop.Animal{Name: "tom"},
		Color:  "blue",
	}
	tom.Show()

	// 多态
	peppa.Hi()
	tom.Hi()
}
