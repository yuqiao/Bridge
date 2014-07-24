package main

import "fmt"
import "time"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	f("direct")
	go f("goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(100 * time.Millisecond)
	fmt.Println("done")
}
