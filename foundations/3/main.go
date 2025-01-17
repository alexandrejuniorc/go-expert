package main

const a = "Hello, World!"

type ID int // ID is a type of int

var (
	b bool    = true
	c int     = 10
	d string  = "Alexandre"
	e float64 = 1.2
	f ID      = 1 // f is a ID type
)

func main() {
	println(f)
}
