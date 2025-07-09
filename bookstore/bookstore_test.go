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
