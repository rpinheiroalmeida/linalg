package util

import (
	"reflect"
	"testing"
)

func TestZip(t *testing.T) {
	cases := []struct {
		v1   []float64
		v2   []float64
		want []Tuple
	}{
		{
			[]float64{1.0, 2.0, 3.0}, []float64{1.0, 3.0, 5.0},
			[]Tuple{Tuple{A: 1.0, B: 1.0}, Tuple{A: 2.0, B: 3.0}, Tuple{A: 3.0, B: 5.0}},
		},
		{
			[]float64{1.0, 2.0}, []float64{1.0, 3.0, 5.0},
			[]Tuple{Tuple{A: 1.0, B: 1.0}, Tuple{A: 2.0, B: 3.0}},
		},
		{
			[]float64{1.0}, []float64{1.0, 3.0, 5.0},
			[]Tuple{Tuple{A: 1.0, B: 1.0}},
		},
		{
			[]float64{1.0, 2.0}, []float64{1.0},
			[]Tuple{Tuple{A: 1.0, B: 1.0}},
		},
	}
	for _, c := range cases {
		gotZip, _ := Zip(c.v1, c.v2)
		if !reflect.DeepEqual(gotZip, c.want) {
			t.Errorf("Zip(%v, %v) want: %v; got: %v", c.v1, c.v2, c.want, gotZip)
		}
	}
}

func TestZip_WhenOneOfParametersIsEmpty(t *testing.T) {
	cases := []struct {
		v1 []float64
		v2 []float64
	}{
		{[]float64{}, []float64{1.0, 3.0, 5.0}},
		{[]float64{1.0, 2.0}, []float64{}},
	}

	for _, c := range cases {
		gotZip, error := Zip(c.v1, c.v2)
		if error == nil {
			t.Errorf("Zip(%v, %v) want: expected error; got %v ", c.v1, c.v2, gotZip)
		}
	}
}
