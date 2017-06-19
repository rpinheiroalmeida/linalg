package vector

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		v1   Vector
		v2   Vector
		want Vector
	}{
		{Vector{1.0, 1.0}, Vector{0.0, 0.0}, Vector{1.0, 1.0}},
		{Vector{1.0, 2.0}, Vector{2.0, 1.0}, Vector{3.0, 3.0}},
	}
	for _, c := range cases {
		got := c.v1.Add(c.v2)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Add(%v, %v) want: %v; got: %v",
				c.v1, c.v2, c.want, got)
		}
	}
}

func TestAdd_WhenVectorsReceiverIsEmpty(t *testing.T) {
	v1, v2 := Vector{}, Vector{}

	defer func() {
		rec := recover()
		if rec == nil {
			t.Errorf("The vector (%v) is empty", v1)
		}
	}()

	v1.Add(v2)
}

func TestAdd_WhenVectorsSenderIsEmpty(t *testing.T) {
	v1, v2 := Vector{1.0}, Vector{}

	defer func() {
		rec := recover()
		if rec == nil {
			t.Errorf("The vector (%v) is empty", v2)
		}
	}()

	v1.Add(v2)
}

func TestAdd_WhenLenVectorsAreDifferent(t *testing.T) {
	v1 := Vector{1.0, 2.0}
	v2 := Vector{1.0}

	defer func() {
		rec := recover()
		if rec == nil {
			t.Errorf("The vectors (%v, %v) have different sizes.", v1, v2)
		}
	}()

	v1.Add(v2)
}

func TestSubtract(t *testing.T) {
	cases := []struct {
		v1   Vector
		v2   Vector
		want Vector
	}{
		{Vector{1.0, 1.0}, Vector{0.0, 0.0}, Vector{1.0, 1.0}},
		{Vector{0.0, 1.0}, Vector{1.0, 2.0}, Vector{-1.0, -1.0}},
	}
	for _, c := range cases {
		got := c.v1.Subtract(c.v2)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Subtract(%v, %v) want: %v; got: %v",
				c.v1, c.v2, c.want, got)
		}
	}
}

func TestSubtract_WhenLenVectorAreDifferent(t *testing.T) {
	v := Vector{1.0, 2.0}
	w := Vector{1.0}

	defer func() {
		if recover() == nil {
			t.Errorf("The vector (%v, %v) have different sizes.", v, w)
		}
	}()

	v.Subtract(w)
}

func TestMax(t *testing.T) {
	cases := []struct {
		Vector
		want float64
	}{
		{Vector{11.0}, 11.0},
		{Vector{11.0, 12.0}, 12.0},
		{Vector{11.0, 20.0, 12.0}, 20.0},
	}

	for _, c := range cases {
		gotMax := c.Max()

		if gotMax != c.want {
			t.Errorf("Expected max (%v) in (%v) but got (%v)", c.want, c.Vector, gotMax)
		}
	}
}

func TestMaxFail(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Expected panic when empty sample")
		}
	}()

	Vector{}.Max()
}

func TestMin(t *testing.T) {
	cases := []struct {
		Vector
		want float64
	}{
		{Vector{13.0}, 13.0},
		{Vector{12.0, 13.0}, 12.0},
		{Vector{12.0, 11.0, 13.0}, 11.0},
	}
	for _, c := range cases {
		gotMin := c.Min()
		if gotMin != c.want {
			t.Errorf("Expected max (%v) in (%v) but got (%v)", c.want, c.Vector, gotMin)
		}
	}
}

func TestMinFail(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Expected panic when empty sample")
		}
	}()

	Vector{}.Min()
}
