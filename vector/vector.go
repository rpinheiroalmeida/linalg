package vector

import (
	"math"

	"github.com/rpinheiroalmeida/linalg/util"
)

type Vector []float64
type binaryCondition func(v1, v2 float64) float64

func validate(v Vector) {
	if len(v) == 0 {
		panic("empty sample supplyed")
	}
}

//Add subtracts two vectors element wise
func (v Vector) Add(w Vector) Vector {
	tuples := util.Zip(v, w)

	result := make(Vector, len(tuples))

	for i, value := range tuples {
		result[i] = value.A + value.B
	}
	return result
}

//Subtract subtracts two vectors elementwise
func (v Vector) Subtract(w Vector) Vector {
	tuples := util.Zip(v, w)

	result := make(Vector, len(tuples))

	for i, value := range tuples {
		result[i] = value.A - value.B
	}
	return result
}

func (v Vector) Dot(w Vector) float64 {
	tuples := util.Zip(v, w)
	var sum float64
	for _, tuple := range tuples {
		sum += tuple.A * tuple.B
	}
	return sum
}

func (v Vector) SumOfSquares() float64 {
	return v.Dot(v)
}

func (v Vector) SquaredDistance(w Vector) float64 {
	result := v.Subtract(w)
	return result.SumOfSquares()
}

func (v Vector) Distance(w Vector) float64 {
	return math.Sqrt(v.SquaredDistance(w))
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
