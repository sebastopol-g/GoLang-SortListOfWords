package input

import (
	"bufio"
	"fmt"
	"os"
)

func ReadInputs() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var words []string
	numberOfWords := 0

	for numberOfWords < 5 {
		fmt.Println("--- Enter your word ---")
		scanner.Scan()
		words = append(words, scanner.Text())
		numberOfWords++
	}

	return words
}
