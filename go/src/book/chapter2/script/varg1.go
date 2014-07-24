package main

import (
	"fmt"
)

func MyPrint(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value")
		case string:
			fmt.Println(arg, "is a string value")
		case int64:
			fmt.Println(arg, "is an int64 value")
		default:
			fmt.Println(arg, "is an unkwown type.")
		}
	}
}

func main() {
	v1 := 1
	var v2 int64 = 234
	v3 := "hello"
	v4 := 1.234
	MyPrint(v1, v2, v3, v4)
}
