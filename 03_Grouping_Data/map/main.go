package main

import (
	"fmt"
)

func main() {
	x := map[string]int{
		"Jasvinder": 30,
		"Jasprit":   29,
		"Parvinder": 29,
		"Lovy":      28,
	}
	fmt.Println(x)
	fmt.Println(x["Lovy"])

	//delete(m map[Key]Type, key Key)
	delete(x, "Jasprit")
	fmt.Println(x)
}
