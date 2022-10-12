package stack

import "testing"

func Test_canCreateStack(t *testing.T) {
	s := NewStack()
	want := true
	got := s.isEmpty()
	if want != got {
		t.Errorf("want %v, but got %v", want, got)
	}
}
