package main

import (
	"fmt"
	"reflect"
)

var (
	name, course, module = "Pravin", "Docker Deep Dive", 3.2
)

func main() {
	fmt.Println("Name is", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("Module is", module, "and is of type", reflect.TypeOf(module))
	a := 10.000000
	b := 3
	c := a + float64(b)
	fmt.Println("\n C has value:", c, "and is of type", reflect.TypeOf(c))
}
