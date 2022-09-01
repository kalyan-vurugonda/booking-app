package main

import "strings"

// validate user input function
func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	// firstName and lastName validation: Needs to enter at least 2 characters.
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	// email validation: Needs to enter valid email format, containing "@" sign.
	isValidEmail := strings.Contains(email, "@")
	// userTickets validation: Needs to enter correct number of tickets (positive + greater than 0).
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
