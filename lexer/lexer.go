package lexer

import "github.com/Jonak-Adipta-Kalita/JAK-MewMew-Language/token"

type Lexer struct {
	input          string
	position       int
	toReadPosition int
	ch             byte
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}

	return l
}

func (l *Lexer) readChar() {
	if l.toReadPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.toReadPosition]
	}

	l.position = l.toReadPosition
	l.toReadPosition += 1
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
