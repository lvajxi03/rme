package utils

import (
	"math/rand/v2"
)

func Wylosuj[T int | float64](min, max T) T {
	if min > max {
		min, max = max, min
	}

	switch minv := any(min).(type) {
	case int:
		maxv := any(max).(int)
		return T(rand.IntN(maxv-minv+1) + minv)
	case float64:
		maxv := any(max).(float64)
		return T(minv + rand.Float64()*(maxv-minv))
	}
	var zero T
	return zero
}
