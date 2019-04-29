// Package leap : all about leap year
package leap

// IsLeapYear to check the input is leap year or not.
func IsLeapYear(year int) bool {
	// Algorithm:
	// if (year is not divisible by 4) then (it is a common year)
	// else if (year is not divisible by 100) then (it is a leap year)
	// else if (year is not divisible by 400) then (it is a common year)
	// else (it is a leap year)

	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
