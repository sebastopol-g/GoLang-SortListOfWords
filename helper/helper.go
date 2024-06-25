package helper

import (
	"fmt"
	"sort"
)

func SortSlice(words []string) []string {

	sort.Strings(words)
	fmt.Printf("Sorted slice is : %v\n", words)
	return words
}
