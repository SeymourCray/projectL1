package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h *Human) sayHello() {
	fmt.Printf("Hello, my name is %s\n", h.name)
}

func (h *Human) sayAge() {
	fmt.Printf("I am %v\n", h.age)
}

// встраивание методов в структуре Action от родительской структуры Human
type Action struct {
	Human
}

func main() {
	human := Human{"Bob", 15}

	action := Action{human}

	action.sayHello()
	action.sayAge()
}
