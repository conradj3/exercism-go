package techpalace

import (
	"strings"
)

const (
	WelcomeMsg        = "Welcome to the Tech Palace, "
	WelcomeBorderChar = "*"
	Newline           = "\n"
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

// Using strings.Repeat
func AddBorderByRepeat(welcomeMsg string, numStarsPerLine int) string {
	return strings.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat("*", numStarsPerLine)
}

// Using strings.Join
func AddBorderByArray(welcomeMsg string, numStarsPerLine int) string {
	var welcomeBorder string
	for i := 0; i < numStarsPerLine; i++ {
		welcomeBorder = strings.Join([]string{welcomeBorder, WelcomeBorderChar}, "")
	}
	message := []string{welcomeBorder + Newline, welcomeMsg + Newline, welcomeBorder}
	return strings.Join(message, "")
}

// Using strings.Builder
func AddBorderByBuilder(welcomeMsg string, numStarsPerLine int) string {
	var sb strings.Builder
	var welcomeBorder string
	for i := 0; i < numStarsPerLine; i++ {
		welcomeBorder = strings.Join([]string{welcomeBorder, WelcomeBorderChar}, "")
	}
	sb.WriteString(welcomeBorder + Newline + welcomeMsg + Newline + welcomeBorder)
	return sb.String()
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	replacedOldMessage := strings.Replace(oldMsg, "*", "", -1)
	replacedOldMessage = strings.TrimSpace(replacedOldMessage)
	return replacedOldMessage
}
