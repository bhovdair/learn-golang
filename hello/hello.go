package main

import (
	"fmt"

	"example/greetings"
)

func main() {
	// fmt.Println(quote.Opt())
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
