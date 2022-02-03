package helper

import "strings"

func IsValidInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) bool {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidNumber := userTickets > 0 && userTickets <= remainingTickets
	isValid := isValidEmail && isValidName && isValidNumber
	return isValid
}
