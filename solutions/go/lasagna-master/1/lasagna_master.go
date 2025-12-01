package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPreparationTime int) int {
    if avgPreparationTime == 0 {
        return len(layers) * 2
    }
    return len(layers) * avgPreparationTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	resNoodles := 0
    resSauce := 0.0
    for _, v := range layers {
        if v == "noodles" {
            resNoodles += 50
        }
        if v == "sauce" {
            resSauce += 0.2
        }
    }
    return resNoodles, resSauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
    myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, num int) []float64 {
    res := make([]float64, len(amounts))
    for i := 0; i < len(amounts); i++ {
        res[i] = amounts[i] * float64(num) / 2.0
    }
    return res
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
