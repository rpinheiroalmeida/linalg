package util

type Tuple struct {
	A, B float64
}

func Zip(a, b []float64) []Tuple {

	if len(a) == 0 || len(b) == 0 {
		return []Tuple{}
	}

	if len(a) < len(b) {
		return createTuples(a, b)
	}
	return createTuples(a[:len(b)], b)

}

func createTuples(a, b []float64) []Tuple {
	tuples := make([]Tuple, len(a), len(a))

	for i, e := range a {
		tuples[i] = Tuple{e, b[i]}
	}
	return tuples

}
