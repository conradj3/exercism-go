package lasagna

// OvenTime returns the amount of time lasagna should be in the oven
const OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	return 2 * numberOfLayers
}

// ElapsedTime calculates the total time needed to create and bake a lasagna.
func ElapsedTime(numberOfLayers int, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}
