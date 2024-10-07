package thefarm

import (
	"errors"
	"fmt"
)

// SillyNephewError struct helps our newphew.
type SillyNephewError struct {
	details int
}

// Error sets a custom error for our silly nephew.
func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.details)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	f, err := weightFodder.FodderAmount()
	// Solve Nephew problem with custom error.
	if cows < 0 {
		return 0, &SillyNephewError{
			details: cows,
		}
	}
	if err == ErrScaleMalfunction && f > 0 {
		return (f * 2 / float64(cows)), nil // Handle scale returns with ErrScaleMalfuction and cows.
	} else if err == ErrScaleMalfunction && f < 0 {
		return 0, errors.New("negative fodder") // Handle negative Fodder with scale error.
	} else if err != nil {
		return 0, err // Handle generic scale error.
	}
	if f < 0 {
		return 0, errors.New("negative fodder") // Handle negative fodder error.
	}

	if cows == 0 {
		return 0, errors.New("division by zero") // Handle division by zero cows, not possible.
	}
	return f / float64(cows), nil
}
