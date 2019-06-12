package main

import (
	"fmt"
)

func main() {
	jb := []string{"Jasvinder", "Singh", "Aulakh", "From", "Ladduwala"}
	fmt.Println(jb)
	lb := []string{"Mango", "banana", "apple", "orange", "pineapple"}
	fmt.Println(lb)
	// append(slice []Type, elems ..Type)

	ln := [][]string{jb, lb}
	fmt.Println(ln)
}
