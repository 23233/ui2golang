package driver

import (
	"math/rand"
	"time"
)

// GetRandomIntInRange returns a random integer between the specified min and max values (inclusive).
// Parameters:
//   - min: The minimum value of the range.
//   - max: The maximum value of the range.
//
// Returns:
//   - A random integer between min (inclusive) and max (exclusive).
func GetRandomIntInRange(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return r.Intn(max-min) + min
}

// GetRandomFloatInRange returns a random floating-point number between the specified min and max values (inclusive).
// Parameters:
//   - min: The minimum value of the range.
//   - max: The maximum value of the range.
//
// Returns:
//   - A random float32 between min (inclusive) and max (exclusive).
func GetRandomFloatInRange(min, max float32) float32 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return r.Float32()*(max-min) + min
}
