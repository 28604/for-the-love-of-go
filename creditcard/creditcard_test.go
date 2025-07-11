package creditcard_test

import (
	"testing"

	"github.com/28604/for-the-love-of-go/creditcard"
)

func TestNew(t *testing.T) {
	t.Parallel()
	want := "1234"
	cc, err := creditcard.New(want)
	if err != nil {
		t.Fatal(err)
	}
	got := cc.Number()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestNewInvalid(t *testing.T) {
	t.Parallel()
	_, err := creditcard.New("")
	if err == nil {
		t.Fatal("Want error for empty credit card number field, got nil.")
	}
}
