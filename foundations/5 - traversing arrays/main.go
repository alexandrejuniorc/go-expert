package main

import "fmt"

func main() {
	var meuArray [3]int // array of 3 integers
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	fmt.Println("Last index of the array:", len(meuArray)-1)
	fmt.Println("Last element of the array:", meuArray[len(meuArray)-1])
	fmt.Println("First element of the array:", meuArray[0])

	// Traversing the array
	for index, value := range meuArray {
		// %d is used to format integers
		// %v is used to format string value
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
