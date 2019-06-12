package main

import (
	"fmt"
)

// slice allows you to group together VALUES of the same type
func main() {
	// x := type{values} //   composite literal
	x := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(len(x))
	fmt.Println(x)
	fmt.Println(x[0])
	fmt.Println(x[4])
	fmt.Println(x[9])
	for i, v := range x {
		fmt.Println(i, v)

	}
}
