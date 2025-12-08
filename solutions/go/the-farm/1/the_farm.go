package thefarm

import (
	"errors"
	"fmt"
)

// type FodderCalculator interface {
// 	FodderAmount(int) (float64, error)
// 	FatteningFactor() (float64, error)
// }

// var errDeterminingAmount = errors.New("amount could not be determined")
// var errDeterminingFactor = errors.New("factor could not be determined")

// TODO: define the 'DivideFood' function
func DivideFood(fodCal FodderCalculator, numCows int) (float64, error) {
	fodderAmount, err := fodCal.FodderAmount(numCows)
	if err != nil {
		return 0, errDeterminingAmount
	}

	fatFactor, err := fodCal.FatteningFactor()
	if err != nil {
		return 0, errDeterminingFactor
	}

	return fodderAmount * fatFactor / float64(numCows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fodCal FodderCalculator, numCows int) (float64, error) {
	if numCows <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fodCal, numCows)
}

type InvalidCowsError struct {
	numCows int
}

func (e *InvalidCowsError) Error() string {
	if e.numCows < 0 {
		return fmt.Sprintf("%d cows are invalid: there are no negative cows", e.numCows)
	}

	return fmt.Sprintf("%d cows are invalid: no cows don't need food", e.numCows)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(numCows int) error {
	if numCows > 0 {
		return nil
	}
	return &InvalidCowsError{
		numCows: numCows,
	}
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
