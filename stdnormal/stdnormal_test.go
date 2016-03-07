package stdnormal

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkInverseTransform(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < b.N; i++ {
		InverseTransform(r)
		//result := InverseTransform(r)
		//_ = result
	}
}

func BenchmarkAcceptReject(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < b.N; i++ {
		AcceptReject(r)
	}
}

func BenchmarkSpecialProperties(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < b.N; i++ {
		Polar(r)
	}
}

func BenchmarkAcceptRejectCounter(b *testing.B) {
	if testing.Short() {
		b.Skip("Running in short mode.")
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	sumTries := 0
	for i := 0; i < b.N; i++ {
		AcceptReject(r)
		_, counter := AcceptReject(r)
		sumTries += counter
	}
	mean := float64(sumTries) / float64(b.N)
	fmt.Printf("Average number of tries: %v\n", mean)
}
