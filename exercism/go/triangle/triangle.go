// Package triangle : all about triangle operation.
package triangle

import (
	"math"
)

// Kind is to determine what kind triangle it is.
type Kind int

const (
	// NaT means not a triangle
	NaT = iota
	// Equ means equilateral triangle
	Equ
	// Iso means isosceles triangle
	Iso
	// Sca means scalene triangle
	Sca
)

// KindFromSides to determine triangle kind by three side.
func KindFromSides(a, b, c float64) Kind {
	if a+b+c == 0 || math.IsInf(a+b+c, 0) || math.IsNaN(a+b+c) || a+b < c || a+c < b || b+c < a {
		return NaT
	}
	if a == b && a == c {
		return Equ
	}
	if a == b || a == c || b == c {
		return Iso
	}
	return Sca
}
