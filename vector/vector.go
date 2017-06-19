package vector

import "math"

type Vector []float64
type binaryCondition func(v1, v2 float64) float64

func validate(v Vector) {
	if len(v) == 0 {
		panic("empty sample supplyed")
	}
}

//Add subtracts two vectors element wise
func (v Vector) Add(w Vector) Vector {
	validate(v)
	validate(w)

	if len(v) != len(w) {
		panic("The vectors have different sizes.")
	}

	result := make([]float64, len(v))

	for i, x := range v {
		result[i] = x + w[i]
	}

	return result
}

//Subtract subtracts two vectors elementwise
func (v Vector) Subtract(w Vector) Vector {
	if len(v) != len(w) {
		panic("The vectors have different sizes.")
	}
	result := make([]float64, len(v))

	for i, x := range v {
		result[i] = x - w[i]
	}

	return result
}

func matchingValue(fn binaryCondition, initial float64, vector Vector) float64 {
	validate(vector)

	current := initial
	for _, value := range vector {
		current = fn(current, value)
	}
	return current
}

func (v Vector) Max() float64 {
	return matchingValue(math.Max, math.Inf(-1), v)
}

func (v Vector) Min() float64 {
	return matchingValue(math.Min, math.Inf(+1), v)
}

func (v Vector) Size() int {
	return len(v)
}

func (v Vector) Empty() bool {
	return v.Size() == 0
}
