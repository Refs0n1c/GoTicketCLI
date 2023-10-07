package main

import "fmt"

func main() {
	var personName string
	var requestedTickets int
	var numberOfAvailTicket = 50
	remainingTickets := numberOfAvailTicket
	welcomeMessage := fmt.Sprintf("Welcome %v, Feel free to book your ticket. There are %v tickets left", personName, remainingTickets)
	fmt.Println(welcomeMessage)
	fmt.Println("Get your tickets NOW!")
	fmt.Println("What is your name?")
	fmt.Scan(&personName)
	fmt.Printf("Wag1 %v ! How many tickets you after? \n", personName)
	fmt.Scan(&requestedTickets)
	fmt.Print("Booking your tickets...")

}
