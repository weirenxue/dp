package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Element interface {
	Value() int
}

type Integer struct {
	value int
}

func NewInteger(value int) *Integer {
	return &Integer{value: value}
}

func (i *Integer) Value() int {
	return i.value
}

type Operator int

const (
	Addition Operator = iota
	Substraction
)

type BinaryOperator struct {
	Type        Operator
	Left, Right Element
}

func (b *BinaryOperator) Value() int {
	switch b.Type {
	case Addition:
		return b.Left.Value() + b.Right.Value()
	case Substraction:
		return b.Left.Value() - b.Right.Value()
	default:
		panic("Unsupported operator")
	}
}

type TokenType int

const (
	Int TokenType = iota
	Plus
	Minus
	Lparen
	Rparen
)

type Token struct {
	Type TokenType
	Text string
}

func Lex(input string) []Token {
	var result []Token

	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '+':
			result = append(result, Token{Type: Plus, Text: "+"})
		case '-':
			result = append(result, Token{Type: Minus, Text: "-"})
		case '(':
			result = append(result, Token{Type: Lparen, Text: "("})
		case ')':
			result = append(result, Token{Type: Rparen, Text: ")"})
		default:
			sb := strings.Builder{}
			j := i
			for ; j < len(input); j++ {
				if unicode.IsDigit(rune(input[j])) {
					sb.WriteRune(rune(input[j]))
				} else {
					break
				}
			}
			i = j - 1
			result = append(result, Token{Type: Int, Text: sb.String()})
		}
	}

	return result
}

func Parse(tokens []Token) Element {
	result := BinaryOperator{}
	haveLhs := false

	for i := 0; i < len(tokens); i++ {
		token := &tokens[i]
		switch token.Type {
		case Int:
			n, _ := strconv.Atoi(token.Text)
			integer := NewInteger(n)
			if !haveLhs {
				result.Left = integer
				haveLhs = true
			} else {
				result.Right = integer
			}
		case Plus:
			result.Type = Addition
		case Minus:
			result.Type = Substraction
		case Lparen:
			j := i
			for ; j < len(tokens); j++ {
				if tokens[j].Type == Rparen {
					break
				}
			}
			element := Parse(tokens[i+1 : j])
			if !haveLhs {
				result.Left = element
				haveLhs = true
			} else {
				result.Right = element
			}
			i = j
		}
	}
	return &result
}

func main() {
	input := "(13+4)-(12+1)"
	tokens := Lex(input)
	fmt.Println(tokens)

	parsed := Parse(tokens)
	fmt.Printf("%s=%d\n", input, parsed.Value())
}
