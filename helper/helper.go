package helper

import (
	"strings"
)

func InputValidation(firstName string, lastName string, email string, userTicket int, remainingTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 1
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := userTicket > 0 && userTicket <= remainingTickets
	return isValidName, isValidEmail, isValidTicket
}
