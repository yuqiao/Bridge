package main

import (
	"fmt"
)

func main() {
	myArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	mySlice := myArray[:5]

	fmt.Println("Elements of myArray: ")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}
	fmt.Println("\nElements of mySlce: ")
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	fmt.Println()

	mySlice2 := make([]int, 5, 10)
	fmt.Println("Len(myslice2):", len(mySlice2))
	fmt.Println("cap(myslice2):", cap(mySlice2))

	mySlice3 := []int{8, 9, 10}
	mySlice2 = append(mySlice2, mySlice3...)
	fmt.Println("\nElements of mySlce2: ")
	for _, v := range mySlice2 {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
