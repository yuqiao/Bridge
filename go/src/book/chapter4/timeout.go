package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1 * time.Second)
		timeout <- true
	}()

	tt := false
	select {
	case <-ch:
		// fmt.Println("Read ch:", <-ch)
	case tt <- timeout:
		fmt.Println("timeout:", tt)
	}
	fmt.Println("finished")
}
