// To check the coverage of tests (coverage of the system code)
// $ go test -cover

// To see what lines are covered and what aren't.
// $ go test -coverprofile=coverage.out
// $ go tool cover -html=coverage.out

// Make sure the coverage is as high as possible before moving on.
// Check if every behavior is correct.

package bookstore_test

import (
	"sort"
	"testing"

	"github.com/28604/for-the-love-of-go/bookstore"
	// To use cmp package, we must update the module with:
	// $ go mod tidy
	"github.com/google/go-cmp/cmp"
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

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := map[int]bookstore.Book{
		1: {ID: 1, Title: "Title 1"},
		2: {ID: 2, Title: "Title 2"},
	}
	want := []bookstore.Book{
		{ID: 1, Title: "Title 1"},
		{ID: 2, Title: "Title 2"},
	}
	got := bookstore.GetAllBooks(catalog)
	// sort.Slice sorts the given slice according to the provided less function.
	// In this case, got is sorted in ascending order for the ID field.
	sort.Slice(got, func(i, j int) bool {
		// Though it is valid to refer to fields of map using got[i].ID,
		// it is invalid to assign values to it, e.g. got[i].ID = 1 (won't work).
		return got[i].ID < got[j].ID
	})
	// Equal compares slices and other data types that cannot be compare with "==".
	// Diff shows the difference between the two objects.
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := map[int]bookstore.Book{
		1: {ID: 1, Title: "Title 1"},
		2: {ID: 2, Title: "Title 2"},
	}
	want := bookstore.Book{ID: 2, Title: "Title 2"}
	got, err := bookstore.GetBook(catalog, 2)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBookBadID(t *testing.T) {
	t.Parallel()
	catalog := map[int]bookstore.Book{}
	_, err := bookstore.GetBook(catalog, 999)
	if err == nil {
		t.Error("Want error for non-existent ID, got nil.")
	}
}
