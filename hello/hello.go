package main

import (
	"fmt"
	"log"
	"github.com/tipdaddy78/go_learning/tree/master/greetings"
)

func main() {
	//Set properties of logger including
	// the log entry prefix
	// flag to diable printing
	// time, source file, and line number
	log.SetPrefix("greetings: ")
	logs.SetFlags(0)

	//Request a message
    message, err := greetings.Hello("")
	//If an error was returned print it to console and exist
	if err != nil {
		log.Fatal(err)
	}

	//If no error, print message to console.
	fmt.Println(message)	
}