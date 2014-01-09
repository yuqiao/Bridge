package main

import "fmt"

type person struct {
	name string
	age  int
}

func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func main() {
	var tom person

	tom.name, tom.age = "Tome", 18

	bob := person{age: 25, name: "Bob"}
	//	paul := person{"Paul", 43}

	tb_Older, tb_diff := Older(tom, bob)

	fmt.Printf("Of %s and %s, %s is older by %d years\n", tom.name, bob.name, tb_Older.name, tb_diff)
}
