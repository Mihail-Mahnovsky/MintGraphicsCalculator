package main

import "math"

type MathParser struct {
	pos    int
	tokens []Token
	x      float64
}

func (m *MathParser) Eval(tokens []Token, x float64) float64 {
	m.tokens = tokens
	m.pos = 0
	m.x = x

	return m.expression()

	//return Graphic{}
}

func (m *MathParser) current() Token {
	if m.pos >= len(m.tokens) {
		return Token{t: Eof}
	}
	return m.tokens[m.pos]
}

func (m *MathParser) eat(t int) {
	if m.current().t == t {
		m.pos++
	}
}

func (m *MathParser) expression() float64 {
	return m.addSub()
}

func (m *MathParser) addSub() float64 {
	value := m.mulDiv()

	for {
		switch m.current().t {
		case Plus:
			m.eat(Plus)
			value += m.mulDiv()
		case Minus:
			m.eat(Minus)
			value -= m.mulDiv()
		default:
			return value
		}
	}
}

func (m *MathParser) mulDiv() float64 {
	value := m.factor()

	for {
		switch m.current().t {
		case Mul:
			m.eat(Mul)
			value *= m.factor()
		case Div:
			m.eat(Div)
			value /= m.factor()
		default:
			return value
		}
	}
}

func (m *MathParser) factor() float64 {
	token := m.current()

	switch token.t {
	case X:
		m.eat(X)
		return m.x
	case Number:
		m.eat(Number)
		return token.v
	case Pi:
		m.eat(Pi)
		return math.Pi
	case Gr:
		m.eat(Gr)
		return math.Phi
	case LParen:
		m.eat(LParen)
		value := m.expression()
		m.eat(RParen)
		return value
	case Minus:
		m.eat(Minus)
		return -m.factor()
	}

	return 0
}
