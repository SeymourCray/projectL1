package main

import "fmt"

type computer interface {
	connectUSB()
}

type newLaptop struct {
}

func (o *newLaptop) connectUSBTypeC() {
	fmt.Println("connected USB type C...")
}

type oldLaptop struct {
}

func (o *oldLaptop) connectUSB() {
	fmt.Println("connected USB...")
}

type person struct {
}

func (p *person) insertUSB(c computer) {
	fmt.Println("inserted USB...")
	c.connectUSB()
}

type adapter struct {
	laptop *newLaptop
}

func (a *adapter) connectUSB() {
	a.laptop.connectUSBTypeC()
}

func main() {
	p := person{}

	o := &oldLaptop{}

	p.insertUSB(o)

	n := &newLaptop{}
	//придерживаемся того же интерфейса, который ожидает person
	a := &adapter{n}

	p.insertUSB(a)
}
