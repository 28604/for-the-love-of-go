package bookstore

// To use the struct Book outside of the "bookstore" package,
// we must capitalised the first letter of the struct "Book".
type Book struct {
	Title  string
	Author string
	Copies int
}
