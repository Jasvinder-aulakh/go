package main

import (
	"fmt"
)

func main() {
	x := []int{11, 22, 33, 44, 55, 66, 77}
	fmt.Println(x)
	//append(slice []Type, elems ..Type)
	x = append(x, 88, 99, 111)
	fmt.Println(x)

	y := []int{222, 333, 444, 555, 666}
	x = append(x, y...)
	fmt.Println(x)
	// delete(m map[Key]Type, key Key)

	x = append(x[:2], x[5:]...)
	fmt.Println(x)

}
