package main

import (
	"strconv"
	"unicode"
)

const (
	Number = iota
	Plus
	Minus
	Mul
	Div
	LParen
	RParen
	Cos
	Sin
	Tan
	Pi
	Gr
	X
	Assign
	Pow
	Sqrt
	Log
	Eof
)

type Token struct {
	t int
	v float64
}

func MakeTokens(line string) []Token {
	tokens := []Token{}

	for pos := 0; pos < len(line); pos++ {
		switch line[pos] {
		case '+':
			tokens = append(tokens, Token{Plus, 0})
		case '-':
			tokens = append(tokens, Token{Minus, 0})
		case '*':
			tokens = append(tokens, Token{Mul, 0})
		case '/':
			tokens = append(tokens, Token{Div, 0})
		case '(':
			tokens = append(tokens, Token{LParen, 0})
		case ')':
			tokens = append(tokens, Token{RParen, 0})
		case 'x':
			tokens = append(tokens, Token{X, 0})
		case '=':
			tokens = append(tokens, Token{Assign, 0})
		case '^':
			tokens = append(tokens, Token{Pow, 0})
		default:
			{
				if unicode.IsLetter(rune(line[pos])) {
					start := pos

					for pos < len(line) && unicode.IsLetter(rune(line[pos])) {
						pos++
					}

					word := line[start:pos]
					pos--

					switch word {
					case "PI":
						tokens = append(tokens, Token{Pi, 0})
					case "GR":
						tokens = append(tokens, Token{Gr, 0})
					case "sin":
						tokens = append(tokens, Token{Sin, 0})
					case "cos":
						tokens = append(tokens, Token{Cos, 0})
					case "tan":
						tokens = append(tokens, Token{Tan, 0})
					case "sqrt":
						tokens = append(tokens, Token{Sqrt, 0})
					case "log":
						tokens = append(tokens, Token{Log, 0})
					}
				} else if unicode.IsDigit(rune(line[pos])) {

					start := pos

					for pos < len(line) &&
						(unicode.IsDigit(rune(line[pos])) || line[pos] == '.') {
						pos++
					}

					numberStr := line[start:pos]
					pos--

					val, _ := strconv.ParseFloat(numberStr, 64)
					tokens = append(tokens, Token{Number, val})
				}
			}
		}

	}

	tokens = append(tokens, Token{Eof, 0})
	return tokens
}
