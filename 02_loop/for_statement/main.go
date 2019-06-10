package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%T\n\t%d\t%v\n\t%#X\n", i, i, i, i)

	}
	for j := 0; j <= 20; j++ {
		fmt.Printf("%T\n\t\t%d\n", j, j)

	}
	for z := 0; z <= 30; z++ {
		if z%2 != 0 {
			continue
		}
		fmt.Println(z)
	}
	for {
		fmt.Println("loop")
		break
	}

	for x := 11; x <= 99; x++ {
		if x%3 != 2 {
			continue
		}
		fmt.Println(x)

	}

}
