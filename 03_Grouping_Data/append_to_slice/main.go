package main

import (
	"fmt"
)

func main() {
	x := []int{5, 10, 88, 100, 239, 237, 500}
	fmt.Println(x)
	// append(slice []Type, elems ..Type)
	x = append(x, 600, 700, 800)
	fmt.Println(x)

	y := []int{3333, 4444, 5555, 6666, 7777}
	fmt.Println(y)
	//append(slice []Type, elems ..Type)
	x = append(x, y...)
	fmt.Println(x)

}
