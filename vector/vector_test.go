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
		{Vector{1.0, 2.0}, Vector{2.0}, Vector{3.0}},
		{Vector{2.0}, Vector{1.0, 2.0}, Vector{3.0}},
		{Vector{}, Vector{}, Vector{}},
		{Vector{2.0}, Vector{}, Vector{}},
		{Vector{}, Vector{2.0}, Vector{}},
	}
	for _, c := range cases {
		got := c.v1.Add(c.v2)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Add(%v, %v) want: %v; got: %v",
				c.v1, c.v2, c.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	cases := []struct {
		v1   Vector
		v2   Vector
		want Vector
	}{
		{Vector{1.0, 1.0}, Vector{0.0, 0.0}, Vector{1.0, 1.0}},
		{Vector{0.0, 1.0}, Vector{1.0, 2.0}, Vector{-1.0, -1.0}},
		{Vector{1.0, 2.0}, Vector{1.0}, Vector{0.0}},
		{Vector{1.0}, Vector{2.0, 3.0}, Vector{-1.0}},
		{Vector{}, Vector{}, Vector{}},
		{Vector{2.0}, Vector{}, Vector{}},
		{Vector{}, Vector{2.0}, Vector{}},
	}
	for _, c := range cases {
		got := c.v1.Subtract(c.v2)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Subtract(%v, %v) want: %v; got: %v",
				c.v1, c.v2, c.want, got)
		}
	}
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

func TestDot(t *testing.T) {
	cases := []struct {
		v1   Vector
		v2   Vector
		want float64
	}{
		{Vector{1.0, 1.0}, Vector{1.0, 1.0}, 2.0},
		{Vector{1.0, 2.0, 3.0}, Vector{4.0, 5.0}, 14.0},
		{Vector{1.0}, Vector{2.0, 5.0}, 2.0},
		{Vector{}, Vector{2.0, 5.0}, 0.0},
	}

	for _, c := range cases {
		gotDot := c.v1.Dot(c.v2)
		if gotDot != c.want {
			t.Errorf("%v.Dot(%v) want: %v but got: %v", c.v1, c.v2, c.want, gotDot)
		}
	}
}

func TestSumOfSquares(t *testing.T) {
	cases := []struct {
		Vector
		want float64
	}{
		{Vector{1.0, 1.0}, 2},
		{Vector{1.0, 2.0, 3.0}, 14.0},
		{Vector{0.0, 0.0, 0.0}, 0.0},
	}

	for _, c := range cases {
		gotSumOfSquares := c.SumOfSquares()
		if c.want != gotSumOfSquares {
			t.Errorf("(%v).SumOfSquares want: %v but got: %v", c.Vector, c.want,
				gotSumOfSquares)
		}
	}
}

// def squared_distance(v, w):
//     return sum_of_squares(vector_subtract(v, w))
func TestSquaredDistance(t *testing.T) {
	cases := []struct {
		v1   Vector
		v2   Vector
		want float64
	}{
		{Vector{1, 1}, Vector{2, 2}, 2},
	}

	for _, c := range cases {
		gotSquaredDistance := c.v1.SquaredDistance(c.v2)
		if c.want != gotSquaredDistance {
			t.Errorf("(%v).SquaredDistance(%v) want: %v but expected: %v", c.v1, c.v2,
				c.want, gotSquaredDistance)
		}
	}
}

func TestDistance(t *testing.T) {
	cases := []struct {
		v1   Vector
		v2   Vector
		want float64
	}{
		{Vector{1.0, 1.0}, Vector{2.0, 2.0}, 1.4142135623730951},
		{Vector{1.0}, Vector{2.0, 2.0}, 1.0},
	}

	for _, c := range cases {
		gotDistance := c.v1.Distance(c.v2)
		if c.want != gotDistance {
			t.Errorf("(%v).Distance(%v) want: %v but got %v", c.v1, c.v2, c.want,
				gotDistance)
		}
	}
}

func TestEmpty(t *testing.T) {
	cases := []struct {
		Vector
		want bool
	}{
		{Vector{1.0}, false},
		{Vector{}, true},
	}

	for _, c := range cases {
		got := c.Empty()
		if got != c.want {
			t.Errorf("(%v).Empty() want: %v but got: %v",
				c.Vector, c.want, got)
		}
	}
}

func TestLess(t *testing.T) {
	cases := []struct {
		Vector
		i, j int
		want bool
	}{
		{Vector{1.0, 2.0}, 0, 1, true},
		{Vector{1.0, 2.0}, 1, 0, false},
		{Vector{1.0, 1.0}, 0, 1, false},
		{Vector{1.0, 1.0}, 1, 0, false},
		{Vector{1.0}, 0, 0, false},
	}

	for _, c := range cases {
		got := c.Less(c.i, c.j)
		if got != c.want {
			t.Errorf("(%v).Less(%d, %d) want: %v but got: %v",
				c.Vector, c.i, c.j, c.want, got)
		}
	}
}

func TestLess_WhenTheValuesAreGreaterThanVector(t *testing.T) {
	cases := []struct {
		Vector
		i, j int
	}{
		{Vector{}, 0, 1},
		{Vector{1.0}, 0, 2},
	}
	for _, c := range cases {
		defer func() {
			if recover() == nil {
				t.Error("Expected panic when the parameters are greater than Vector size")
			}
		}()

		c.Less(c.i, c.j)
	}
}

func TestSwap(t *testing.T) {
	cases := []struct {
		before Vector
		i, j   int
		want   Vector
	}{
		{Vector{1.0, 2.0, 3.0}, 0, 2, Vector{3.0, 2.0, 1.0}},
		{Vector{1.0, 2.0}, 1, 0, Vector{2.0, 1.0}},
		{Vector{1.0, 1.0}, 0, 1, Vector{1.0, 1.0}},
		{Vector{1.0}, 0, 0, Vector{1.0}},
	}

	for _, c := range cases {
		c.before.Swap(c.i, c.j)
		if !reflect.DeepEqual(c.before, c.want) {
			t.Errorf("(%v).Swap(%d, %d) want: %v; got: %v",
				c.want, c.i, c.j, c.want, c.before)
		}
	}
}

func TestSwap_WhenTheValuesAreGreaterThanVector(t *testing.T) {
	cases := []struct {
		Vector
		i, j int
	}{
		{Vector{1.0, 2.0, 3.0}, 0, 4},
		{Vector{}, 0, 0},
	}

	for _, c := range cases {
		defer func() {
			if recover() == nil {
				t.Error("Expected panic when the parameters are greater than Vector size")
			}
		}()
		c.Swap(c.i, c.j)
	}
}
