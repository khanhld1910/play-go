package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Sneeze Cheese")
	fmt.Println(message)
}