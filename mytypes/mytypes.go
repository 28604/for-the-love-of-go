package mytypes

// MyInt is a custom version of `int` type.
type MyInt int

// MyString is a custom version of `string` type.
type MyString string

// Twice returns twice the value of the input MyInt.
func (a MyInt) Twice() MyInt {
	return 2 * a
}

// Len returns the length of the input MyString.
func (s MyString) Len() int {
	return len(s) // Can just use the generic len function.
}
