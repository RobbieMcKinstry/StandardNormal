package stdnormal

import (
	"math"
	"math/rand"
)

var isCached = false
var cache float64

func Polar(r *rand.Rand) float64 {

	var (
		u      = 1.0
		v1, v2 float64
	)

	if isCached {
		isCached = false
		return cache
	}

	for u >= 1.0 { // repeat while u >= 1
		v1 = 2.0 * (r.Float64() - 0.5) // draw a random number from [-1.0, 1.0)
		v2 = 2.0 * (r.Float64() - 0.5) // draw a second number from [-1.0, 1.0)
		u = v1*v1 + v2*v2
	}

	insideRoot := -2.0 * math.Log(u) / u
	angle := math.Sqrt(insideRoot)

	z1 := v1 * angle
	z2 := v2 * angle

	isCached = true
	cache = z2

	return z1
}
