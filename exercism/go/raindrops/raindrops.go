// Package raindrops is a package about factorization.
package raindrops

import (
	"strconv"
	"strings"
)

var rainTypes = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

// Convert a number to a string, the contents of which depend on the number's factors.
// If the number has 3 as a factor, output 'Pling'.
// If the number has 5 as a factor, output 'Plang'.
// If the number has 7 as a factor, output 'Plong'.
// If the number does not have 3, 5, or 7 as a factor, just pass the number's digits straight through.
func Convert(n int) string {
	if n < 1 {
		return ""
	}

	result := ""
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			if i == 3 || i == 5 || i == 7 {
				if !strings.Contains(result, rainTypes[i]) {
					result = result + rainTypes[i]
				}
			}
		}
	}

	if result == "" {
		return strconv.Itoa(n)
	} else {
		return result
	}
}
