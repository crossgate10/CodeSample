// twofer is short for two for one. "One for X, one for me." When X is a name or "you".
package twofer

import (
	"bytes"
	"fmt"
	"strings"
)

// Note: Before go version 1.10 use "bytes.Buffer", after 1.10 use "strings.Builder"
func ShareWith(name string) string {
	if len(strings.TrimSpace(name)) <= 0 {
		return "One for you, one for me."
	}

	ss := []string{
		"One for ",
		name,
		", one for me.",
	}

	var b bytes.Buffer
	for _, s := range ss {
		fmt.Fprintf(&b, s)
	}

	return b.String()
}
