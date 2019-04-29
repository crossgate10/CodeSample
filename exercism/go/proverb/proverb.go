// Package proverb is a demo of the rhyme "For Want of a Nail"
package proverb

import (
	"fmt"
)

// Proverb print the rhyme "For Want of a Nail" by input
func Proverb(rhyme []string) []string {
	format1st := "For want of a %s the %s was lost."
	format2sd := "And all for the want of a %s."

	if len(rhyme) == 0 {
		return []string(nil)
	}

	if len(rhyme) == 1 {
		return []string{fmt.Sprintf(format2sd, rhyme[0])}
	}

	result := make([]string, 0, len(rhyme))
	for i, s := range rhyme {
		if i == len(rhyme)-1 {
			result = append(result, fmt.Sprintf(format2sd, rhyme[0]))
		} else {
			result = append(result, fmt.Sprintf(format1st, s, rhyme[i+1]))
		}
	}
	return result
}
