package mytypes_test

import (
	"testing"

	"github.com/28604/for-the-love-of-go/mytypes"
)

func TestMyIntTwice(t *testing.T) {
	t.Parallel()
	a := mytypes.MyInt(9)
	want := mytypes.MyInt(18)
	got := a.Twice()
	if want != got {
		t.Errorf("Twice(%d): want %d got %d", a, want, got)
	}
}

func TestMyStringLen(t *testing.T) {
	t.Parallel()
	s := mytypes.MyString("Hello, World!")
	want := 13
	got := s.Len()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
