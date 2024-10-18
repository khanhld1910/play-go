package main

import (
	"fmt"
	"time"
)

func main() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// if without else
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// if with a statement before condition
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else {
		fmt.Println(num, "is positive")
	}

	// if with multiple conditions
	if num := 10; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// basic switch
	switch light := "green"; light {
	case "red":
		fmt.Println("Stop")
	case "yellow":
		fmt.Println("Slow down")
	case "green":
		fmt.Println("Go")
	}

	// case with multiple values, and default
	switch dayOfWeek := "huhu"; dayOfWeek {
	case "monday", "wednesday", "friday":
		fmt.Println("Run")
	case "tuesday", "thursday":
		fmt.Println("Swim")
	case "saturday", "sunday":
		fmt.Println("Rest")
	default:
		fmt.Println("Invalid day")
	}

	// switch without condition
	// this behaves different to JavaScript's switch-case

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}

	// note that in Go, only 1 case will be executed, unlike JavaScript it can fall through multiple cases if no break is used

	// can use swtich to check type
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			// print: value is a boolean: i
			fmt.Printf("Value is a boolean: %v\n", t)
		case int:
			// print: value is an int: i
			fmt.Printf("Value is an int: %v\n", t)
		case string:
			// print: value is a string: i
			fmt.Printf("Value is a string: %v\n", t)
		default:
			// print: don't know the type of i
			fmt.Printf("Don't know the type of: %v\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
	whatAmI([]int{1, 2})
}
