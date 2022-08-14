package main

import (
	"fmt"
	"log"

	"sh9h.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	// A slice of names
	names := []string{"Sh9h", "Mummas", "Tamanna"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)

}
