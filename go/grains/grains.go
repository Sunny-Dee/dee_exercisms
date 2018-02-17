package grains

import "errors"

// Square calculates the total grains of rice on the specified square
func Square(n int) (uint64, error) {
	var v uint64 = 1

	if n < 1 || n > 64 {
		return 0, errors.New("Invalid Square")
	}

	return v << (uint)(n-1), nil
}

// Total gives the grain total for all 64 squares (1 to 64)
func Total() uint64 {
	var sum uint64
	for i := 1; i <= 64; i++ {
		v, _ := Square(i)
		sum += v
	}
	return sum
}
