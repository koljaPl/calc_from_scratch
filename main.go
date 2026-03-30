package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Test the provided example first
	testExample()

	fmt.Println("\n=== Go Calculator ===")
	fmt.Println("Enter mathematical expressions or 'quit' to exit")
	fmt.Println("Supported operations: +, -, *, /, ^")
	fmt.Println("Supported functions: log, ln, sin, cos, tan, sqrt, abs")
	fmt.Println("Example: 15 + 13 * 43 - (log(8) - 1)")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">> ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "quit" || input == "exit" {
			break
		}

		if input == "" {
			continue
		}

		result, err := Evaluate(input)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("= %f\n", result)
		}
		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
	}

	fmt.Println("Goodbye!")
}
