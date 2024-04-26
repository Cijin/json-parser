package main

import "fmt"

type lexer struct {
	input []byte

	ch byte

	// current lexer position
	position int

	// position to start reading next ch from
	readPosition int
}

func NewLexer(input []byte) *lexer {
	l := &lexer{input: input}
	l.readChar()

	return l
}

func (l *lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // 0 -> null
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *lexer) peak() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func (l *lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *lexer) readString() {
	// read next char as current is "
	l.readChar()
	// keep reading till " encountered
	for l.ch != '"' {
		// reached end
		if l.ch == 0 {
			return
		}

		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *lexer) NextChar() error {
	l.skipWhiteSpace()

	switch l.ch {
	case '"':
		l.readString()
		if l.ch == 0 {
			return fmt.Errorf("unterminated string")
		}

	case ',':
		if l.peak() != '\n' {
			return fmt.Errorf("trailing comma")
		}

	default:
		if isLetter(l.ch) {
			return fmt.Errorf("invalid character")
		}
	}

	l.readChar()
	return nil
}
