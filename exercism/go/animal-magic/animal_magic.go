package chance

import (
	"math/rand"
	"time"
)

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	return rand.Intn(20) + 1
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	// min + rand.Float64() * (max - min)
	return rand.Float64() * 12.0
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	rand.Seed(time.Now().UnixNano())

	a := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	f := func(i, j int) {
		a[i], a[j] = a[j], a[i]
	}
	rand.Shuffle(len(a), f)
	return a
}
