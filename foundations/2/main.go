package main

const a = "Hello, World!"

// global variables
var (
	b bool    = true
	c int     = 10
	d string  = "Alexandre"
	e float64 = 1.2
)

func main() {
	// b = 10 // cannot assign to b (declared const) - cannot change the value of a constant

	// local variable
	// var a string = "X"
	a := 'X' // short declaration operator
	a = 'Y'
	println(a)
}

func x() {}
