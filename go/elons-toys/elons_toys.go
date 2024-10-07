package elon

import "fmt"

// Drive updates the number of meters driven based on the car's speed, and reduces the battery according to the battery drainage
func (car *Car) Drive() {
	if car.battery >= car.batteryDrain {
		car.distance += car.speed
		car.battery -= car.batteryDrain
	}
}

// DisplayDistance returns the distance as displayed on the LED display
func (car *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// DisplayBattery returns the battery percentage as displayed on the LED display
func (car *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

// CanFinish checks if a car is able to finish the race by set trackDistance
func (car *Car) CanFinish(distance int) bool {
	return distance <= (car.battery / car.batteryDrain * car.speed)
}
