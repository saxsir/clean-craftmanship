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

func Test_afterOnePush_isNotEmpty(t *testing.T) {
	s := NewStack()
	s.push(0)
	want := false
	got := s.isEmpty()
	if want != got {
		t.Errorf("want %v, but got %v", want, got)
	}
}

func Test_afterOnePushAndOnePop_isEmpty(t *testing.T) {
	s := NewStack()
	s.push(0)
	s.pop()
	want := true
	got := s.isEmpty()
	if want != got {
		t.Errorf("want %v, but got %v", want, got)
	}
}
