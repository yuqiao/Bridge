package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 22})
	fmt.Println(person{age: 32, name: "Alice"})
	fmt.Println(person{name: "Peter"})
	fmt.Println(&person{name: "Ann", age: 42})

	s := person{"Sean", 50}
	fmt.Println("name:", s.name)

	sp := &s
	fmt.Println("age:", sp.age)

	sp.age = 51
	fmt.Println(s.age)
}
