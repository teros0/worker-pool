package main

import (
	"fmt"
)

type A struct {
	F    func(...interface{}) interface{}
	Args []interface{}
}

func NewA(f func(...interface{}), args ...interface{}) A {
	j := A{F: f,
		Args: args}
	return j
}

func a(x, y int) {
	fmt.Println(x + y)
}

func main() {
	fmt.Println("TO ENTER THE POOOL PUT YOUR SHOES OFF")
	b := a
	qwe := NewA(a, 1, 2, 3)
	fmt.Printf("%d", qwe.F(qwe.Args...))
}
