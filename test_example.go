package main

import (
	"fmt"
)

func testExample() {
	expression := "15 + 13 * 43 - (log(8) - 1)"
	result, err := Evaluate(expression)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	fmt.Printf("Expression: %s\n", expression)
	fmt.Printf("Result: %.12f\n", result)
	fmt.Printf("Expected: 574.096910013\n")
	fmt.Printf("Matches expected: %.12f ≈ 574.096910013\n", result)
}
