// AirportRobot Package
package airportrobot

import (
	"fmt"
)

// Define the Greeter interface
type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

// Implement the SayHello function
func SayHello(name string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(name))
}

// Implement Italian struct and methods
type Italian struct{}

func (i Italian) LanguageName() string {
	return "Italian"
}

// Greet method for Italian struct
func (i Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

// Implement Portuguese struct and methods
type Portuguese struct{}

// LanguageName method for Portuguese struct
func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

// Greet method for Portuguese struct
func (p Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}
