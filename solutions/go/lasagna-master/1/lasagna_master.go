package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timePerLayer int) int {
    if timePerLayer == 0 {
        timePerLayer = 2
    }

    return len(layers) * timePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
    for _, layer := range layers {
        switch layer {
            case "noodles":
        		noodles += 50
            case "sauce":
        		sauce += 0.2
        }
    }
    return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsRecipe, myRecipe []string) {
    myRecipe[len(myRecipe) - 1] = friendsRecipe[len(friendsRecipe) - 1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(portions []float64, servings int) []float64 {
    scale := float64(servings) / 2
    scaled := make([]float64, len(portions))
    for i, portion := range portions {
        scaled[i] = portion * scale
    }
	return scaled
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
