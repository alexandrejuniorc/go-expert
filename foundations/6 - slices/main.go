package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Printf("Length:%d Capacity:%d %v\n", len(slice), cap(slice), slice)
	fmt.Printf("Length:%d Capacity:%d %v\n", len(slice[:0]), cap(slice[:0]), slice[:0]) // shows items from index 0 to 0
	fmt.Printf("Length:%d Capacity:%d %v\n", len(slice[:4]), cap(slice[:4]), slice[:4]) // shows items from index 0 to 3
	fmt.Printf("Length:%d Capacity:%d %v\n", len(slice[2:]), cap(slice[2:]), slice[2:]) // shows items from index 2 to 8

	slice = append(slice, 110)                                                          // append 12 to the slice
	fmt.Printf("Length:%d Capacity:%d %v\n", len(slice[:2]), cap(slice[:2]), slice[:2]) // shows items from index 0 to 1
	fmt.Printf("Length:%d Capacity:%d %v\n", len(slice), cap(slice), slice)             // shows all items
}
