package main

import "fmt"

const a = "Hello, World!"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Alexandre"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	// %T refers to the type of the variable in Go-syntax representation
	fmt.Printf("O tipo de E é %T\n", e)

	// #v refers to the value of the variable in a Go-syntax representation
	fmt.Printf("O valor de E é %v\n", e)

	// More example
	fmt.Printf("O tipo de F é %T\n", f)
}
