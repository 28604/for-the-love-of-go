package mytypes_test

import (
	"strings"
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

func TestStringsBuilder(t *testing.T) {
	t.Parallel()
	var sb strings.Builder
	sb.WriteString("Hello, ")
	sb.WriteString("Gopher!")

	want := "Hello, Gopher!"
	got := sb.String()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}

	wantLen := 14
	gotLen := sb.Len()
	if wantLen != gotLen {
		t.Errorf("want %d, got %d", wantLen, gotLen)
	}
}

func TestMyStringsBuilder(t *testing.T) {
	t.Parallel()
	var mb mytypes.MyBuilder
	mb.Contents.WriteString("Hello, ")
	mb.Contents.WriteString("Gopher!")

	want := "Hello, Gopher!"
	got := mb.Contents.String()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}

	wantLen := 14
	gotLen := mb.Contents.Len()
	if wantLen != gotLen {
		t.Errorf("want %d got %d", wantLen, gotLen)
	}
}

func TestStringUpperCaser(t *testing.T) {
	t.Parallel()
	var su mytypes.StringUppercaser
	su.Contents.WriteString("Hello, Gopher!")
	want := "HELLO, GOPHER!"
	got := su.ToUpper()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestDouble(t *testing.T) {
	t.Parallel()
	var x int = 2
	want := 4
	// Parameters are assed by value,
	// so Double(x) will not change x if passing in an int.
	// &x is of type pointer to int.
	mytypes.Double(&x)
	if want != x {
		t.Errorf("want %d got %d", want, x)
	}
}

func TestMyIntDouble(t *testing.T) {
	t.Parallel()
	x := mytypes.MyInt(12)
	want := mytypes.MyInt(24)
	(&x).MyIntDouble()
	if want != x {
		t.Errorf("want %d, got %d", want, x)
	}
}
