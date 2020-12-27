// Package iteration exercises iteration
package iteration

import (
	"strings"
)

// Repeat repeats the given character by declared number of times
func Repeat(character string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}

func standardRepeat(character string, count int) string {
	return strings.Repeat(character, count)
}
