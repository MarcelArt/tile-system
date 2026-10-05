package rng

import "math/rand/v2"

func Float32(min, max float32) float32 {
	rng := rand.Float32()
	return min + rng*(max-min)
}

func Int32(min, max int32) int32 {
	rng := rand.N(max - min)
	return min + rng
}
