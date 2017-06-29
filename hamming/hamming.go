package hamming

import "errors"

const testVersion = 6

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("String length does not match.")
	}
	var counter int
	for i := 0; i < len(a); i++ {
		if string(a[i]) != string(b[i]) {
			counter++
		}
	}
	return counter, nil
}
