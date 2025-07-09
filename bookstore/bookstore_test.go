// To check the coverage of tests (coverage of the system code)
// $ go test -cover

// To see what lines are covered and what aren't.
// $ go test -coverprofile=coverage.out
// $ go tool cover -html=coverage.out

// Make sure the coverage is as high as possible before moving on.
// Check if every behavior is correct.

package bookstore_test

import (
	"testing"

	"github.com/28604/for-the-love-of-go/bookstore"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Condo",
		Copies: 2,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Title 1",
		Author: "John Doe",
		Copies: 3,
	}
	want := 2
	result, err := bookstore.Buy(b)
	if err != nil {
		t.Fatal(err)
	}
	got := result.Copies
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestBuyErrorsIfNoCopiesLeft(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Title 2",
		Author: "James Doe",
		Copies: 0,
	}
	_, err := bookstore.Buy(b)
	if err == nil {
		t.Error("Want error buying from zero copies, got nil.")
	}
}
