package stdnormal

import (
	"math"
	"math/rand"
)

// y = (1 - ln x) / -1.702
func InverseTransform(r *rand.Rand) float64 {
	x := r.Float64()               // generate a random number between [0 , 1)
	numerator := 1.0 - math.Log(x) // calculate the numerator
	denominator := -1.702          // calculate the numerator

	return CoinFlip(numerator/denominator, r) // calculate the displacement from the origin
}

// This function flips the sign of the input 50% of the time.
// This is used when only positive/negative values for x are generated, and thus it generates the other half of the normal line.
func CoinFlip(x float64, r *rand.Rand) float64 {

	flip := r.Float64() // used to ensure that 50% of numbers of negative, since the normal curve is not bijective
	if flip < 0.5 {
		x *= -1.0
	}
	return x
}
