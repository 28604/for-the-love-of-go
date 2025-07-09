package bookstore

import "errors"

// To use the struct Book outside of the "bookstore" package,
// we must capitalised the first letter of the struct "Book".
type Book struct {
	Title  string
	Author string
	Copies int
}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}
