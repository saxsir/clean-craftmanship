package prime_factors

import "testing"
import "golang.org/x/exp/slices"

func isEmpty(s []int) bool {
	return len(s) == 0
}

func contains(a, b []int) bool {
	return slices.Compare(a, b) == 0
}

func Test_factors(t *testing.T) {
	cases := []struct {
		in  int
		out []int
	}{
		{in: 2, out: []int{2}},
		{in: 3, out: []int{3}},
		{in: 4, out: []int{2, 2}},
		{in: 5, out: []int{5}},
		{in: 6, out: []int{2, 3}},
		{in: 7, out: []int{7}},
		{in: 8, out: []int{2, 2, 2}},
		{in: 9, out: []int{3, 3}},
	}

	want := true
	if got := isEmpty(factorsOf(1)); want != got {
		t.Errorf("want %v, but got %v", want, got)
	}

	for _, c := range cases {
		if got := contains(factorsOf(c.in), c.out); want != got {
			t.Errorf("case %v: want %v, but got %v", c, want, got)
		}
	}
}

func factorsOf(n int) []int {
	s := make([]int, 0)

	for divisor := 2; n > 1; divisor++ {
		for ; n%divisor == 0; n /= divisor {
			s = append(s, divisor)
		}
	}

	return s
}
