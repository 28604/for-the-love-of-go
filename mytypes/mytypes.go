package mytypes

import "strings"

// When creating a new type, it does not inherit the methods.
// To use the methods of the underlying type,
// we could define a struct and add the underlying type as a field.

// MyInt is a custom version of `int` type.
type MyInt int

// MyString is a custom version of `string` type.
type MyString string

// MyBuilderCannotUseMethod is a self-defined type of strings.Builder,
// but cannot use the methods of strings.Builder.
// type MyBuilderCannotUseMethod strings.Builder

// Unless we create a struct to wrap strings.Builder.
type MyBuilder struct {
	Contents strings.Builder
}

type StringUppercaser struct {
	Contents strings.Builder
}

// Twice returns twice the value of the input MyInt.
func (a MyInt) Twice() MyInt {
	return 2 * a
}

// Len returns the length of the input MyString.
func (s MyString) Len() int {
	return len(s) // Can just use the generic len function.
}

func (mb MyBuilder) Hello() string {
	return "Hello, Gopher!"
}

func (su StringUppercaser) ToUpper() string {
	// Because ToUpper only works for strings (also we want to return string),
	// so we must convert to string for uppercasing.
	return strings.ToUpper(su.Contents.String())
}

func Double(input *int) {
	// input must be dereferenced to be multiplied by 2 (an integer).
	*input *= 2
}

func (input *MyInt) MyIntDouble() {
	*input *= MyInt(2)
}
