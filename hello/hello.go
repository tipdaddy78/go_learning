package main

import (
	"fmt"
	"../greetings"
)

func main() {
    message := greetings.Hello("Tipper")
	fmt.Println(message)	
}