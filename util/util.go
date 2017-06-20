package util

import (
	"fmt"
)

type Tuple struct {
	A, B float64
}

func Zip(a, b []float64) ([]Tuple, error) {

	if len(a) == 0 || len(b) == 0 {
		return nil, fmt.Errorf("Zip: one of the arguments is empty")
	}

	if len(a) < len(b) {
		return createTuples(a, b), nil
	}
	return createTuples(a[:len(b)], b), nil

}

func createTuples(a, b []float64) []Tuple {
	tuples := make([]Tuple, len(a), len(a))

	for i, e := range a {
		tuples[i] = Tuple{e, b[i]}
	}
	return tuples

}
