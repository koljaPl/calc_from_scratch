package main

import (
	"fmt"
	"math"
	"strconv"
)

type Parser struct {
	tokens []Token
	pos    int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{tokens: tokens, pos: 0}
}

func (p *Parser) current() Token {
	if p.pos >= len(p.tokens) {
		return Token{TokenEOF, ""}
	}
	return p.tokens[p.pos]
}

func (p *Parser) advance() {
	p.pos++
}

func (p *Parser) eat(tokenType TokenType) error {
	if p.current().Type != tokenType {
		return fmt.Errorf("expected %v, got %v", tokenType, p.current().Type)
	}
	p.advance()
	return nil
}

func (p *Parser) parse() (float64, error) {
	result, err := p.parseExpression()
	if err != nil {
		return 0, err
	}

	if p.current().Type != TokenEOF {
		return 0, fmt.Errorf("unexpected token: %v", p.current())
	}

	return result, nil
}

func (p *Parser) parseExpression() (float64, error) {
	return p.parseBinaryOp(1)
}

func (p *Parser) parseBinaryOp(minPrecedence int) (float64, error) {
	left, err := p.parseUnary()
	if err != nil {
		return 0, err
	}

	for p.current().Type == TokenOperator {
		op := p.current().Value
		precedence := getPrecedence(op)

		if precedence < minPrecedence {
			break
		}

		p.advance()

		var right float64
		if isLeftAssociative(op) {
			right, err = p.parseBinaryOp(precedence + 1)
		} else {
			right, err = p.parseBinaryOp(precedence)
		}

		if err != nil {
			return 0, err
		}

		left = applyOperator(op, left, right)
	}

	return left, nil
}

func (p *Parser) parseUnary() (float64, error) {
	if p.current().Type == TokenOperator && (p.current().Value == "-" || p.current().Value == "+") {
		op := p.current().Value
		p.advance()

		operand, err := p.parseUnary()
		if err != nil {
			return 0, err
		}

		if op == "-" {
			return -operand, nil
		}
		return operand, nil
	}

	return p.parsePrimary()
}

func (p *Parser) parsePrimary() (float64, error) {
	if p.current().Type == TokenNumber {
		value := p.current().Value
		p.advance()

		num, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid number: %s", value)
		}

		return num, nil
	}

	if p.current().Type == TokenFunction {
		fn := p.current().Value
		p.advance()

		if err := p.eat(TokenLParen); err != nil {
			return 0, err
		}

		arg, err := p.parseExpression()
		if err != nil {
			return 0, err
		}

		if err := p.eat(TokenRParen); err != nil {
			return 0, err
		}

		if fn == "log" && arg <= 0 {
			return 0, fmt.Errorf("log of non-positive number")
		}
		if fn == "ln" && arg <= 0 {
			return 0, fmt.Errorf("ln of non-positive number")
		}
		if fn == "sqrt" && arg < 0 {
			return 0, fmt.Errorf("sqrt of negative number")
		}

		return applyFunction(fn, arg), nil
	}

	if p.current().Type == TokenLParen {
		p.advance()

		value, err := p.parseExpression()
		if err != nil {
			return 0, err
		}

		if err := p.eat(TokenRParen); err != nil {
			return 0, err
		}

		return value, nil
	}

	return 0, fmt.Errorf("unexpected token: %v", p.current())
}

func Evaluate(expression string) (float64, error) {
	tokens, err := Tokenize(expression)
	if err != nil {
		return 0, err
	}

	parser := NewParser(tokens)
	result, err := parser.parse()
	if err != nil {
		return 0, err
	}

	if math.IsNaN(result) {
		return 0, fmt.Errorf("result is NaN")
	}

	if math.IsInf(result, 0) {
		return 0, fmt.Errorf("result is infinite")
	}

	return result, nil
}
