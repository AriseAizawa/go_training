package main

import "fmt"

type person struct {
	name string
	sex  string
}

type greeter interface {
	greet()
}

func (p person) greet() {
	if p.sex == "Male" {
		fmt.Print("Hello,Mr.")
		fmt.Println(p.name)
	} else {
		fmt.Print("Hello,Ms.")
		fmt.Println(p.name)
	}

}

func main() {
	p01 := person{
		name: "Aizawa",
		sex:  "Male",
	}

	p01.greet()
}
