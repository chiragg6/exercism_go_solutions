package raindrops

import "fmt"

// task is to convert a number into string

func Convert(n int) string {
	s := ""
	if n%3 == 0 {
		s += "Pling"
	}
	if n%5 == 0 {
		s += "Plang"
	}
	if n%7 == 0 {
		s += "Plong"
	}
	if len(s) > 0 {
		return s
	}
	return fmt.Sprintf("%d", n)
}
