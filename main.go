package main

import (
	"fmt"
)

// below are the package level variables
var confName = "Go Conference"

const confTicket = 50

var remainingTickets = 50

var bookings_slice = make([]UserData, 0) //slice(dynamic)

type UserData struct {
	firstName  string
	lastName   string
	email      string
	userTicket int
}

func main() {

	greetUsers()

	for {
		//get the user input
		firstName, lastName, email, userTicket := getUserInput()

		//validate the inputs given by the user
		isValidName, isValidEmail, isValidTicket := inputValidation(firstName, lastName, email, userTicket, remainingTickets)

		if isValidName && isValidEmail && isValidTicket {
			// book tickets
			bookTicket(firstName, lastName, email, userTicket)

			//printing the first names
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				//end program
				fmt.Printf("Go conference is booked out. Come back next year.\n")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Firstname/Lastname you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("@ character is missing in your mail id.")
			}
			if !isValidTicket {
				fmt.Println("Wrong number of tickets entered")
			}
		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v Booking Application\n", confName)
	fmt.Println("We have a total of", confTicket, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings_slice {
		firstNames = append(firstNames, booking.firstName)

	}
	return firstNames
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTicket int

	//passing the memory address of the variable so that the user entered value is stored in that particular location
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email ID: ")
	fmt.Scan(&email)
	fmt.Print("Enter the number of tickets to be booked: ")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

func bookTicket(firstName string, lastName string, email string, userTicket int) {
	//creating a map for the user
	var userData = UserData{
		firstName:  firstName,
		lastName:   lastName,
		email:      email,
		userTicket: userTicket,
	}
	bookings_slice = append(bookings_slice, userData)
	fmt.Printf("Thank you %v %v with email %v has booked %v tickets.\n", firstName, lastName, email, userTicket)
	remainingTickets = remainingTickets - userTicket
	fmt.Printf("%v tickets are remaining for the %v\n", remainingTickets, confName)
	fmt.Printf("List of Bookings: %v\n", bookings_slice)
}
