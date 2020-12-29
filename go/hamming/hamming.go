package hamming

import "errors"

func Distance(a, b string) (int, error) {
	distance := 0
	if len(a) == len(b) {
		for i := range a {
			if a[i] != b[i] {
				distance++

			}

		}
		return distance, nil

	} else {
		return -1, errors.New("Lenght is not same of the strings")
	}
}
