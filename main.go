package main

import (
	"fmt"
)

// fmt meaning --format package
// it's a compiled language and starts executing from main function.

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("yess we are learning go")
	fmt.Println("Hello, World!")
	fmt.Println("yess we are learning go")

	// println and Println is differnent.
	//

	var conferenv = "conferenece"
	const conferrenceTickets = 50
	var remaningTickets = 50

	fmt.Println("hello heloo", conferenv, ",bello beloo")
	fmt.Println("we have total of this many tickets", remaningTickets)
	fmt.Println("welcome to", conferenv, "my dear friends")
	fmt.Printf("welcome to %v my dear friends\n", conferenv)

	// %v is a placeholder

	var userName string

	// you need to tell the type of the variable
	var userTicket int
	userName = "sai"
	userTicket = 10

	// here the 'Printf' is used to print the string and the variable

	fmt.Printf("User %v booked %v tickets\n", userName, userTicket)

}

// if you initilize the var or import the module and dont use it --you'll get an error
