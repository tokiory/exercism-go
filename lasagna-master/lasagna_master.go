package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, preparationTime int) int {
	if preparationTime == 0 {
		preparationTime = 2
	}

	return len(layers) * preparationTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var noodlesAmount int
	var sauceAmount float64

	for _, layer := range layers {
		if layer == "noodles" {
			noodlesAmount += 50
			continue
		}

		if layer == "sauce" {
			sauceAmount += 0.2
		}
	}

	return noodlesAmount, sauceAmount
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, ownList []string) {
	(ownList)[len(ownList)-1] = (friendList)[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, n int) []float64 {
	recipe := make([]float64, len(quantities))
	for i := range quantities {
		recipe[i] = float64(quantities[i]/2) * float64(n)
	}

	return recipe
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
