package speed

import "math"

// Car implements a remote controlled car.
type Car struct {
	speed        int
	batteryDrain int

	battery  int
	distance int
}

type Track struct {
	distance int
}

// NewCar creates a new car with given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      100,
		batteryDrain: batteryDrain,
		speed:        speed,
		distance:     0,
	}
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	return Track{distance: distance}
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery < car.batteryDrain {
		return car
	}
	return Car{
		speed:        car.speed,
		batteryDrain: car.batteryDrain,
		battery:      car.battery - car.batteryDrain,
		distance:     car.distance + car.speed,
	}
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	td := float64(track.distance)
	carTimes := int(math.Ceil(td / float64(car.speed)))
	carDrain := car.batteryDrain * carTimes
	return carDrain <= car.battery
}
