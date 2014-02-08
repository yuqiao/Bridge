package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, val := range nums {
		sum += val
	}
	fmt.Println("sum:", sum)

	for i, val := range nums {
		if val == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for key, val := range kvs {
		fmt.Printf("%s -> %s \n", key, val)
	}

	for i, c := range "go" {
		fmt.Println(i, ". ", c)
	}
}
