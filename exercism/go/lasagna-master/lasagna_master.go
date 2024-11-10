package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, ptime int) int {
	if ptime == 0 {
		ptime = 2
	}
	return len(layers) * ptime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodle int, sauce float64) {
	for _, v := range layers {
		if v == "noodles" {
			noodle += 50
		}
		if v == "sauce" {
			sauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amount []float64, portion int) []float64 {
	amount2 := make([]float64, len(amount))
	copy(amount2, amount)

	for i, v := range amount {
		amount2[i] = v / 2 * float64(portion)
	}
	return amount2
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
