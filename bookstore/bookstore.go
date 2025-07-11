package bookstore

import (
	"errors"
	"fmt"
)

// To use the struct Book outside of the "bookstore" package,
// we must capitalised the first letter of the struct "Book".
type Book struct {
	ID              int
	Title           string
	Author          string
	Copies          int
	PriceCents      int
	DiscountPercent int
}

type Catalog map[int]Book

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}

// It is a convention to name the receiver "c" short.
func (c Catalog) GetAllBooks() []Book {
	// Cannot directly return catalog because catalog is a mapping.
	// To return the full slice of book, we can append each book using a for loop.
	result := []Book{}
	// Slices are ordered, but maps are random.
	// So, by appending the result with some order, unsorted results between tests might differ.
	// Therefore, we can sort the result during test to ensure results aligned each time.
	for _, b := range c {
		result = append(result, b)
	}
	return result
}

func (c Catalog) GetBook(ID int) (Book, error) {
	// ok is a boolean value showing if there's a mapping of the key.
	b, ok := c[ID]
	if !ok {
		// To show which value is invalid, we must use fmt.Errorf instead of errors.New.
		return Book{}, fmt.Errorf("ID %d does not exist", ID)
	}
	return b, nil
}

// Method is a function of a type of objects.
// Think of method as a dynamic struct field that computes upon calling it.
// A method must be defined in the same package as the type.
// To define a method in another package, we can define a new type and add methods on top of it.
func (b Book) NetPriceCents() int {
	return b.PriceCents * (100 - b.DiscountPercent) / 100
}

// As long as we specify the object type as a pointer,
// Go would auto-dereference the pointer for us even when we call the method
// like b.SetPriceCents(123), rather than (&b).SetPriceCents(123).
func (b *Book) SetPriceCents(price int) error {
	if price < 0 {
		err := fmt.Errorf("negative price %d", price)
		return err
	}
	b.PriceCents = price
	return nil
}
