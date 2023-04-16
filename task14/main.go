package main

import (
	"fmt"
	"reflect"
)

func defineType(v interface{}) string {
	//assertion type
	switch v.(type) {
	case int:
		return "int"
	case bool:
		return "bool"
	case string:
		return "string"
	case chan int:
		return "chan int"
	default:
		return reflect.TypeOf(v).String()
	}
}

func main() {
	a := 1
	b := true
	c := "example"
	d := make(chan int)

	fmt.Println(defineType(a))
	fmt.Println(defineType(b))
	fmt.Println(defineType(c))
	fmt.Println(defineType(d))
}
