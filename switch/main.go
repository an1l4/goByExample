package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	switch i {
	case 1:
		fmt.Println("i is 1")
	case 2:
		fmt.Println("i is 2")
	case 3:
		fmt.Println("i is 3")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("week days")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("before afternoon")
	default:
		fmt.Println("after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("i am boolean")
		case int:
			fmt.Println("i am integer")
		default:
			fmt.Println("i dont know", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
