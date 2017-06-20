package util

import (
	"fmt"
)

type Tuple struct {
	a, b float64
}

func Zip(a, b []float64) ([]Tuple, error) {

	if len(a) == 0 || len(b) == 0 {
		return nil, fmt.Errorf("Zip: one of the arguments is empty")
	}

	r := make([]Tuple, len(a), len(a))

	for i, e := range a {
		r[i] = Tuple{e, b[i]}
	}

	return r, nil
}
