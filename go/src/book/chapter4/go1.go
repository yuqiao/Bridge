package main

import (
	"fmt"
)

/*func Add(x, y int) {
	z := x + y
	fmt.Println(x, "+", y, "=", z)
}*/

func Count(ch chan int) {
	fmt.Println("Counting")
	ch <- 1
}

func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}

	for _, ch := range chs {
		value := <-ch
		fmt.Println("value:", value)
	}
}
