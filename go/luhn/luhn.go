// Package luhn exposes a single method, Valid, which tests for
// validity via the Luhn algorithm
package luhn

import "unicode"

const testVersion = 2

// reverse returns its argument string reversed rune-wise left to right.
// cribbed from:
// https://github.com/golang/example/blob/master/stringutil/reverse.go
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Valid returns true if input passes the Luhn checksum algorithm
// see: https://en.wikipedia.org/wiki/Luhn_algorithm
// validation will fail if string contains any values other than
// digits and spaces.
func Valid(input string) bool {

	sum := 0     // the rolling sum value
	counter := 0 // separate parallel counter to only count digit positions

	// reverse the string to allow idiomatic rune ranging
	// also using our custom counter, so ignore the given index
	for _, r := range reverse(input) {

		// process according to Luhn if the rune at this position is
		// a digit
		if unicode.IsDigit(r) {
			// q & d rune to int, may not work on non-ascii compat runes
			val := int(r - '0')

			// double every *second* digit from the rightmost position
			// (leftmost since this is the reversed input)
			if counter%2 == 1 {
				val = (val * 2)
				// if doubled value is greater than 9 then subtract 9
				// from the value
				if val > 9 {
					val = val - 9
				}
			}
			sum += val

			// only increment counter for valid input positions (digits)
			counter++
			continue
		}

		// ignore spaces entirely
		if unicode.IsSpace(r) {
			continue
		}

		// if we reach this point this rune is not a digit or space,
		// therefore invalid input string.
		return false
	}

	// if the valid inputs are 1 or less it is an invalid string
	if counter < 2 {
		return false
	}

	// valid input, but can it luhn?
	// all we need to is if the sum is evenly dividable by 10
	return (sum % 10) == 0
}
