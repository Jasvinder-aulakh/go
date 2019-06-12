package main

import (
	"fmt"
)

func main() {
	x := []int{1, 3, 5, 7, 9}
	fmt.Println(x[0])
	fmt.Println(x)
	// slicing_a_slice
	fmt.Println(x[:2])
	fmt.Println(x[2:])
	fmt.Println(x[:])
	fmt.Println(x[2:4])

	// just with for
	for i, v := range x {
		fmt.Println(i, v)
	}
	// with range (It will same print)
	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])

	}
}
