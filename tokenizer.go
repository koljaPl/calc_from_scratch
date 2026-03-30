package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

type TokenType int

const (
	TokenNumber TokenType = iota
	TokenOperator
	TokenLParen
	TokenRParen
	TokenFunction
	TokenEOF
)

type Token struct {
	Type  TokenType
	Value string
}

func isOperator(c rune) bool {
	return c == '+' || c == '-' || c == '*' || c == '/' || c == '^'
}

func isFunction(s string) bool {
	functions := []string{"log", "ln", "sin", "cos", "tan", "sqrt", "abs"}
	for _, f := range functions {
		if s == f {
			return true
		}
	}
	return false
}

func Tokenize(input string) ([]Token, error) {
	var tokens []Token
	input = strings.ReplaceAll(input, " ", "")

	i := 0
	for i < len(input) {
		c := rune(input[i])

		if unicode.IsDigit(c) || c == '.' {
			var numStr string
			for i < len(input) && (unicode.IsDigit(rune(input[i])) || input[i] == '.') {
				numStr += string(input[i])
				i++
			}
			tokens = append(tokens, Token{TokenNumber, numStr})
			continue
		}

		if isOperator(c) {
			tokens = append(tokens, Token{TokenOperator, string(c)})
			i++
			continue
		}

		if c == '(' {
			tokens = append(tokens, Token{TokenLParen, string(c)})
			i++
			continue
		}

		if c == ')' {
			tokens = append(tokens, Token{TokenRParen, string(c)})
			i++
			continue
		}

		if unicode.IsLetter(c) {
			var funcStr string
			for i < len(input) && unicode.IsLetter(rune(input[i])) {
				funcStr += string(input[i])
				i++
			}
			if isFunction(funcStr) {
				tokens = append(tokens, Token{TokenFunction, funcStr})
			} else {
				return nil, fmt.Errorf("unknown function: %s", funcStr)
			}
			continue
		}

		return nil, fmt.Errorf("invalid character: %c", c)
	}

	tokens = append(tokens, Token{TokenEOF, ""})
	return tokens, nil
}

func getPrecedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	case "^":
		return 3
	default:
		return 0
	}
}

func isLeftAssociative(op string) bool {
	return op != "^"
}

func applyOperator(op string, a, b float64) float64 {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	case "^":
		return math.Pow(a, b)
	default:
		return 0
	}
}

func applyFunction(fn string, a float64) float64 {
	switch fn {
	case "log":
		return math.Log10(a)
	case "ln":
		return math.Log(a)
	case "sin":
		return math.Sin(a)
	case "cos":
		return math.Cos(a)
	case "tan":
		return math.Tan(a)
	case "sqrt":
		return math.Sqrt(a)
	case "abs":
		return math.Abs(a)
	default:
		return 0
	}
}
