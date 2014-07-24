package main

import (
	"errors"
	"fmt"
	"mymath"
)

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func Add(a int, b int) (ret int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("Should be non-negative numbers")
		return
	}
	return a + b, nil
}

func main() {
	personDB := make(map[string]PersonInfo)
	personDB["1234"] = PersonInfo{"1", "Jack", "Room 101,..."}

	person, ok := personDB["1234"]
	if ok {
		fmt.Println("Found person", person.Name, "with Id 1234.")
	} else {
		fmt.Println("Did not Fid person with Id 1234.")
	}

	ret, err := Add(-1, 3)
	if err == nil {
		fmt.Println("The Result:", ret)
	} else {
		fmt.Println("Add error:", err)
	}

	fmt.Println("plus:", mymath.Plus(1, 2, 3, 4))
}
