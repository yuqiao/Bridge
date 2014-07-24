package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start@", time.Now())
	timer1 := time.NewTimer(time.Second * 2)

	<-timer1.C
	fmt.Println("timer1 expired. @", time.Now())

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer2 expired. @", time.Now())
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer2 stopped. @", time.Now())
	}
}
