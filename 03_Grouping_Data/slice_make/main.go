package main

import (
	"fmt"
)

func main() {
	// make(Type, size IntegerType)
	x := make([]int, 15, 17)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// // XXX: []
	x[0] = 55
	x[5] = 65
	x[14] = 112
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	// append(slice []Type, elems ..Type)

	x = append(x, 999)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	//append(slice []Type, elems ..Type)
	x = append(x, 1111)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	//append(slice []Type, elems ..Type)
	x = append(x, 2222)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println("Have you notices here soething ^")

	//append(slice []Type, elems ..Type)

	x = append(x, 3333)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println("Have you notices here too soething ^(36)")

}
