package lasagna

const (
	noodles                = "noodles"
	sauce                  = "sauce"
	defaultAveragePerLayer = 2
	gramsPerNoodle         = 50
	litersPerSauce         = 0.2
)

// PreparationTime returns the estimate for the total preparation time based on the number of layers as an int.
func PreparationTime(layers []string, averagePerLayer int) int {
	if averagePerLayer == 0 {
		averagePerLayer = defaultAveragePerLayer
	}
	return averagePerLayer * len(layers)
}

// Quantities will return the quantity of noodles and sauce needed to make your meal.
func Quantities(layers []string) (n int, s float64) {
	for i := range layers {
		if layers[i] == noodles {
			n += gramsPerNoodle
		} else if layers[i] == sauce {
			s += litersPerSauce
		}
	}
	return n, s
}

// AddSecretIngredient generates a new slice and add the last item from your friends list to the end of your list. Neither argument should be modified
func AddSecretIngredient(friendsList, myList []string) []string {
	length := len(friendsList)
	if length == 0 {
		return myList
	}
	return append(myList, friendsList[length-1])
}

// ScaleRecipe returns a slice of float64 of the amounts needed for the desired number of portions.
func ScaleRecipe(quantities []float64, portions int) []float64 {
	var multiplier = float64(portions) / 2
	ret := make([]float64, len(quantities))
	for i := range quantities {
		ret[i] = quantities[i] * multiplier
	}
	return ret
}
