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
	if got := s.getSize(); 1 != got {
		t.Errorf("want %v, but got %v", 1, got)
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

	if got := s.getSize(); 0 != got {
		t.Errorf("want %v, but got %v", 1, got)
	}
}

func Test_afterTwoPushes_sizeIsTwo(t *testing.T) {
	s := NewStack()
	s.push(0)
	s.push(0)
	want := 2
	got := s.getSize()
	if want != got {
		t.Errorf("want %v, but got %v", want, got)
	}
}
