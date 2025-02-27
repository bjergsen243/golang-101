package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("basic switch")
	basicSwitch()
	fmt.Println("multiple case switch")
	multipleCaseSwitch()
	fmt.Println("no condition switch")
	noConditionSwitch()
	fmt.Println("switch type")
	switchType()
}

func basicSwitch() {
	i := 2
	fmt.Println("i is", i)
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
}

func multipleCaseSwitch() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
}

func noConditionSwitch() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
}

func switchType() {
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
