package util

import (
	"reflect"
	"testing"

	"github.com/rpinheiroalmeida/linalg/vector"
)

func TestZip(t *testing.T) {
	cases := []struct {
		v1   vector.Vector
		v2   vector.Vector
		want []Tuple
	}{
		{
			vector.Vector{1.0, 2.0, 3.0}, vector.Vector{1.0, 3.0, 5.0},
			[]Tuple{Tuple{a: 1.0, b: 1.0}, Tuple{a: 2.0, b: 3.0}, Tuple{a: 3.0, b: 5.0}},
		},
		{
			vector.Vector{1.0, 2.0}, vector.Vector{1.0, 3.0, 5.0},
			[]Tuple{Tuple{a: 1.0, b: 1.0}, Tuple{a: 2.0, b: 3.0}},
		},
		{
			vector.Vector{1.0}, vector.Vector{1.0, 3.0, 5.0},
			[]Tuple{Tuple{a: 1.0, b: 1.0}},
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
		v1 vector.Vector
		v2 vector.Vector
	}{
		{vector.Vector{}, vector.Vector{1.0, 3.0, 5.0}},
		{vector.Vector{1.0, 2.0}, vector.Vector{}},
	}

	for _, c := range cases {
		gotZip, error := Zip(c.v1, c.v2)
		if error == nil {
			t.Errorf("Zip(%v, %v) want: expected error; got %v ", c.v1, c.v2, gotZip)
		}
	}
}
