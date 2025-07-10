package bookstore

import "errors"

// To use the struct Book outside of the "bookstore" package,
// we must capitalised the first letter of the struct "Book".
type Book struct {
	ID     int
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

func GetAllBooks(catalog []Book) []Book {
	return catalog
}

func GetBook(catalog []Book, ID int) Book {
	for _, b := range catalog {
		if b.ID == ID {
			return b
		}
	}
	return Book{}
}
