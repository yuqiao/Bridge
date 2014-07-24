package main

import (
	"fmt"
)

func main() {
	str := "Hello,世界"
	ch := str[0]
	fmt.Printf("The lenght of \"%s\" is %d \n", str, len(str))
	fmt.Printf("The first character of '%s' is %c. \n", str, ch)

	for i := 0; i < len(str); i++ {
		ch := str[i]
		fmt.Println(i, ch)
	}

	for i, ch := range str {
		fmt.Println(i, ch)
	}
}
