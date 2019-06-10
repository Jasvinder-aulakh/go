package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("Should not print")
	case (2 == 4):
		fmt.Println("This should not print too")
	case (4 == 4):
		fmt.Println("Print it")
		fallthrough
	case (5 == 5):
		fmt.Println("Again this too print")
		fallthrough
	case (7 == 9):
		fmt.Println("Not print")
		fallthrough
	case (13 == 13):
		fmt.Println("Value 13")
	case (13 == 15):
		fmt.Println("Not 15 equal to 13, not to print")
		// fallthrough will print every syntex for next
		fallthrough
	case (17 == 18):
		fmt.Println("This is not equal to 18")
		// this is not printed because we have not used fallthrough continue, when it breaks, it will not read next
	}

	switch "bond" {
	case "money":
		fmt.Println("This is not bond")
	case "bond":
		fmt.Println("This is bond and it get print")
	default:
		fmt.Println("Default Print")

	}
	x := "Jasvinder"
	switch x {
	case "Money":
		fmt.Println("will not print")
	case "Jasvinder":
		fmt.Println("Definetly it will print")
	default:
		fmt.Println("Default Print")

	}

}
