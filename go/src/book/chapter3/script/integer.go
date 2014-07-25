package main

import (
	"fmt"
)

type Integer int

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

type Lesser interface {
	Less(b Integer) bool
}

func (a Integer) Less(b Integer) bool {
	return a < b
}

func (a *Integer) Add(b Integer) {
	*a += b
}

func main() {
	var a Integer = 1
	if a.Less(2) {
		fmt.Println(a, "Less 2")
	}
	a.Add(2)
	fmt.Println("a = ", a)

	var b LessAdder = &a
	var c Lesser = a
	var d Lesser = &a

	b.Add(3)
	fmt.Println("a = ", a)
	fmt.Println("c lesser 4:", c.Less(4))
	fmt.Println("d lesser 4:", d.Less(4))
}
