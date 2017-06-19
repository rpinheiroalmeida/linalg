package vector

type Vector []float64

func (v Vector) validate() {
	if len(v) == 0 {
		panic("empty sample supplyed")
	}
}

//Add subtracts two vectors element wise
func (v Vector) Add(w Vector) Vector {
	v.validate()

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
