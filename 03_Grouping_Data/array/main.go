package main

import (
	"fmt"
)

func main() {
	var x [5]int
	fmt.Println(x)
	x[0] = 10
	x[3] = 40
	fmt.Println(x)
	fmt.Println(len(x))
}
