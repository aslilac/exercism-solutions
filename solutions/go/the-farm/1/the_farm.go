package thefarm

import "fmt"
import "errors"

// TODO: define the 'DivideFood' function
func DivideFood(calc FodderCalculator, cows int) (float64, error) {
    fodder, err := calc.FodderAmount(cows)
    if err != nil {
        return 0, err
    }

    ff, err := calc.FatteningFactor()
    if err != nil {
        return 0, err
    }

    return fodder * ff / float64(cows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(calc FodderCalculator, cows int) (float64, error) {
    if cows <= 0 {
        return 0, errors.New("invalid number of cows")
    }

    return DivideFood(calc, cows)
}

type InvalidCowsError struct {
    cows int
    message string
}

func NewInvalidCowsError(cows int, message string) *InvalidCowsError {
    return &InvalidCowsError { cows, message }
}

func (self InvalidCowsError) Error() string {
    return fmt.Sprintf("%v cows are invalid: %s", self.cows, self.message)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cows int) error {
    if cows == 0 {
        return NewInvalidCowsError(cows, "no cows don't need food")
    }

    if cows < 0 {
        return NewInvalidCowsError(cows, "there are no negative cows")
    }

    return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
