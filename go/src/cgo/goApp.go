package main

// #include "somelibc.h"
import "C"

import (
	"fmt"
)

func main() {
	fmt.Println("1+2=", C.sum(1, 2))
}
