// Package hamming : about hamming distance operations.
// hamming diatance: calculate the difference between two string.
package hamming

import "errors"

// Distance : Get hamming distance.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("wrong input")
	}

	distance := 0
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
