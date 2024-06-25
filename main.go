package main

import (
	"fmt"

	"training.go/order/helper"
	"training.go/order/input"
)

func main() {

	fmt.Println("--- This program allows you to enter 5 words ---")
	fmt.Println("--- It will then sort the words ---")
	words := input.ReadInputs()
	fmt.Printf("---- Words are : %v ---", words)
	sortedSlice := helper.SortSlice(words)
	fmt.Printf("---- Sorted words : %v ---", sortedSlice)
	fmt.Println("--- End of program ---")
}
