package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
		for j := 0; j < 20; j++ {
			fmt.Printf("%d\t%T\t%b\n", i, i, j)

		}

	}

}
