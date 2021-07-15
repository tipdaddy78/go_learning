package main

import (
	"fmt"
	"github.com/tipdaddy78/go_learning/tree/master/greetings"
)

func main() {
    message := greetings.Hello("Tipper")
	fmt.Println(message)	
}