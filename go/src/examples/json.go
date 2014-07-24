package main

import (
	"encoding/json"
	"fmt"
	//"os"
)

type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	slcD := []string{"apple", "orange", "banana"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "banana": 6}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "orange", "banana"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println("res1B:", string(res1B))

	res2D := &Response2{
		Page:   1,
		Fruits: []string{"apple", "orange", "banana"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println("res2B:", string(res2B))
}
