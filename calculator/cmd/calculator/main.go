// To run a Go code, there must exist a "main" package and a "main" function.
// In Go, we cannot have more than one package in the same folder, except for test packages.
package main

import (
	"fmt"

	"github.com/28604/for-the-love-of-go/calculator"
)

func main() {
	result := calculator.Add(2, 2)
	fmt.Println(result)
}

// To run the program (in the calculator parent folder)
// $ go run ./cmd/calculator/main.go

// To produce the executable binary and run the program (in the calculator parent folder)
// $ go build -o anyNameForTheBinary ./cmd/calculator
// $ ./anyNameForTheBinary
