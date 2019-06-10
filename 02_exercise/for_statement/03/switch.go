package main

import (
	"fmt"
	"time"
)

func main() {
	i := 3
	fmt.Println("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its the weekend")
	default:
		fmt.Println("It's a weekday")
	}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("Its after noon")
	}

	whatami := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am a int")
		default:
			fmt.Println("Dont know type %T\n", t)
		}

	}
	whatami(true)
	whatami(1)
	whatami("hey")

}
