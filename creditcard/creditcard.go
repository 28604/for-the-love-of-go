package creditcard

import "errors"

// The "card" type and "number" field are never called in the test
// because we cannot access anything unexported outside the "creditcard" package.
// So, to access the "number" field, we use an accessor method, Number.
// To create a "card", rather than creditcard.card, we use creditcard.New(num).
type card struct {
	number string
}

// New is a "consturctor" (a function that returns valid object).
func New(num string) (card, error) {
	if num == "" {
		return card{}, errors.New("number field cannot be empty")
	}
	return card{number: num}, nil
}

func (c card) Number() string {
	return c.number
}
