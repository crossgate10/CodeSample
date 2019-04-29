// Package strand is a package about strand transcription.
package strand

var rnaTypes = map[string]string{
	"G": "C",
	"C": "G",
	"T": "A",
	"A": "U",
}

// ToRNA given a DNA strand, return its RNA complement (per RNA transcription).
// G -> C
// C -> G
// T -> A
// A -> U
func ToRNA(dna string) string {
	if len(dna) == 1 {
		return rnaTypes[dna]
	}

	result := ""
	for _, c := range dna {
		result = result + rnaTypes[string(c)]
	}
	return result
}
