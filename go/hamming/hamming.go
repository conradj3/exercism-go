package hamming

import "errors"

// Distance calculates the number of different characters between string arguments
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("cannot be unequal size")
	}
	count := 0
	for i, s := range []byte(a) {
		if s != b[i] {
			count++
		}
	}
	return count, nil
}
