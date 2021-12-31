package purchase

import "sort"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	} else {
		return false
	}
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	options := []string{option1, option2}
	sort.Strings(options)
	return options[0] + " is clearly the better choice."

}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age == 0 {
		return originalPrice
	} else if age >= 1 && age <= 3 {
		// Apply 20% discount for age 1-3
		return originalPrice * .80
	} else if age >= 3 && age < 10 {
		// Apply 30% discount for age 3-9
		return originalPrice * .70
	} else {
		// Apply 50% discount age 10+
		return originalPrice * .50
	}
}
