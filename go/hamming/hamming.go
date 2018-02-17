package hamming

import "errors"

// Distance calculates two DNA strings hamming distance
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		err := errors.New("Error DNA strings must be the same length")
		return -1, err
	}

	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}

	return count, nil
}
