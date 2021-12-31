package techpalace

import (
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var loyalityGreeting string
	for i := 0; i < numStarsPerLine; i++ {
		loyalityGreeting += "*"
	}
	displayMessage := loyalityGreeting + "\n" + welcomeMsg + "\n" + loyalityGreeting
	return displayMessage
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	replacedOldMessage := strings.Replace(oldMsg, "*", "", -1)
	replacedOldMessage = strings.TrimSpace(replacedOldMessage)
	return replacedOldMessage
}
