package main

import (
	"fmt"
	"strings"
	"unicode"
)

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

func main() {
	input := "(13+4)-(12+1)"
	tokens := Lex(input)
	fmt.Println(tokens)
}
