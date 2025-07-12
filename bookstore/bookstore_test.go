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
	"github.com/google/go-cmp/cmp/cmpopts"
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
	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "Title 1"},
		2: {ID: 2, Title: "Title 2"},
	}
	want := []bookstore.Book{
		{ID: 1, Title: "Title 1"},
		{ID: 2, Title: "Title 2"},
	}
	got := catalog.GetAllBooks()
	// sort.Slice sorts the given slice according to the provided less function.
	// In this case, got is sorted in ascending order for the ID field.
	sort.Slice(got, func(i, j int) bool {
		// Though it is valid to refer to fields of map using got[i].ID,
		// it is invalid to assign values to it, e.g. got[i].ID = 1 (won't work).
		return got[i].ID < got[j].ID
	})
	// Equal compares slices and other data types that cannot be compare with "==".
	// Diff shows the difference between the two objects.
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "Title 1"},
		2: {ID: 2, Title: "Title 2"},
	}
	want := bookstore.Book{ID: 2, Title: "Title 2"}
	got, err := catalog.GetBook(2)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		// Adding an unexported field makes the TestGetBook fail
		// because cmp cannot access unexported fields.
		// To compare Book, we now can use compopts.IgnoreUnexported
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBookBadID(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{}
	_, err := catalog.GetBook(999)
	if err == nil {
		t.Error("Want error for non-existent ID, got nil.")
	}
}

func TestNetPriceCents(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		PriceCents:      100,
		DiscountPercent: 30,
	}
	want := 70
	got := b.NetPriceCents()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestSetPriceCents(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{ID: 3, PriceCents: 2000}
	want := 1500
	err := b.SetPriceCents(want)
	// This line below does the same thing, but explicitly.
	// err := (&b).SetPriceCents(want)
	if err != nil {
		t.Fatal(err)
	}
	got := b.PriceCents
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestSetPriceCentsBadPrice(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{ID: 3, PriceCents: 2000}
	err := b.SetPriceCents(-1)
	if err == nil {
		t.Fatal("Want err for invalid price, got nil.")
	}
}

func TestSetCategory(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{ID: 1, Title: "Title 1"}
	cats := []bookstore.Category{
		bookstore.CategoryAutobiography,
		bookstore.CategoryLargePrintRomance,
		bookstore.CategoryParticlePhysics,
	}
	for _, cat := range cats {
		err := b.SetCategory(cat)
		if err != nil {
			t.Fatal(err)
		}
		got := b.Category()
		if cat != got {
			t.Errorf("want %q, got %q", cat, got)
		}
	}
}

func TestSetCategoryInvalid(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{ID: 3, Title: "Title 1"}
	err := b.SetCategory(999)
	if err == nil {
		t.Fatal("Want error for invalid category, got nil.")
	}
}
