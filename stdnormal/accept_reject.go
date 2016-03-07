package stdnormal

import (
	"math"
	"math/rand"
)

// pdf := (2 / sqrt(2 * Pi)) * exp(-0.5 * x^2)
// c := sqrt(2e/p)
// f(x) := pdf
// g(x) := exp(-x)
func AcceptReject(randomGen *rand.Rand) (float64, int) {
	counter := 1
	for {
		x := randomGen.ExpFloat64()
		r := randomGen.Float64()
		c := math.Sqrt(math.E / math.Pi)
		if r*c*g(x) <= f(x) {
			x := CoinFlip(x, randomGen)
			return x, counter
		}
		counter++
	}
}

func f(x float64) float64 {
	pi := math.Pi
	temp1 := 2.0 / math.Sqrt(2*pi)
	temp2 := 1.0 + math.Expm1(-0.5*x*x)
	return temp1 * temp2
}

// return exp(x)
func g(x float64) float64 {
	return math.Expm1(-x) + 1.0
}
