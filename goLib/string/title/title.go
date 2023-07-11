// https://pkg.go.dev/strings#Title
// func
// Title
// DEPRECATED
// func Title(s string) string
// Title returns a copy of the string s with all Unicode letters that begin words mapped to their Unicode title case.

// Deprecated: The rule Title uses for word boundaries does not handle Unicode punctuation properly. Use golang.org/x/text/cases instead.

package main

import (
	"fmt"
	"strings"
)

func main() {
	// Compare this example to the ToTitle example.
	fmt.Println(strings.Title("her royal highness"))
	fmt.Println(strings.Title("loud noises"))
	fmt.Println(strings.Title("хлеб"))
}
