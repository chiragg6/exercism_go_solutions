// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"unicode"
)

func Abbreviate(s string) string {
	prior := ' '
	acronym := make([]rune, 0)
	for _, c := range s {
		if prior == ' ' || prior == '-' || prior == '_' {
			if c != ' ' && c != '_' && c != '-' {
				acronym = append(acronym, unicode.ToUpper(c))

			}

		}
		prior = c
	}
	return string(acronym)
}

// func abbrivate(s string) string {
// 	previous := ' '
// 	sep := '_'
// 	acronym := make([]rune, 0)
// 	for _, c := range s {

// 		if c != previous && c != sep {
// 			acronym = append(acronym, c)
// 		} else {
// 			acronym = append(acronym, -1)
// 		}

// 	}
// 	var words []string

// 	str := string(acronym)
// 	ac := make([]string, 50)

// 	words = strings.Split(str, "-1")
// 	for i, word := range words {
// 		ac := append(ac, unicode.ToUpper(word))

// 	}

// }

// func abb(s string) []string {

// 	for i, w := range s {
// 		if w == ' ' || w == '-' {
// 			w = "-1"
// 		}
// 	}

// }

// func splitString(s string, sepraters []string) []string {
// 	var result []string

// 	word := ""
// 	for _, char := range s {
// 		add := true
// 		for _, sep := range sepraters {
// 			if sep == string(char) {
// 				if word != "" {
// 					result = append(result, word)
// 					word = ""
// 					add = false
// 					break
// 				}

// 			}
// 		}
// 		if add {
// 			word += string(char)
// 		}

// 	}
// 	if word != "" {
// 		result = append(result, word)
// 	}
// 	return result
// }
