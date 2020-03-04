package bob

import (
	"strings"
	"unicode"
)

const TestVersion = 1

func Hey(greeting string) string {
	greeting = strings.TrimSpace(greeting)

	switch {
	case greeting == "":
		return "Fine. Be that way"
	case any(greeting, unicode.IsUpper) && !any(greeting, unicode.IsLower):
		return "Calm down, I know what I'm doing!"
	case greeting[len(greeting)-1] == '?':
		return "Sure."
	default:
		return "Whatever."
	}

}

func any(greet string, test func(rune) bool) bool {
	for _, item := range greet {
		if test(item) {
			return true
		}
	}
	return false
}
