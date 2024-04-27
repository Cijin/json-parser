package main

import (
	"bytes"
	"fmt"
)

var keywords = [][]byte{
	[]byte("null"),
	[]byte("true"),
	[]byte("false"),
}

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
	for isWhiteSpace(l.ch) {
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

func (l *lexer) readArray() error {
	// currently at '['
	l.readChar()

	for l.ch != ']' {
		if l.ch == 0 {
			return fmt.Errorf("array not terminated")
		}

		if err := l.NextChar(); err != nil {
			return err
		}
	}

	// currently at ']'
	l.readChar()

	return nil
}

// returns keyword and advances read position
func (l *lexer) readKeyword() error {
	var keyword []byte

	for l.ch != ',' || isWhiteSpace(l.ch) {
		if l.ch == 0 {
			return fmt.Errorf("unexpected EOF")
		}

		keyword = append(keyword, l.ch)

		l.readChar()
	}

	if !isValidKeyword(keyword) {
		return fmt.Errorf("unidentified keyword")
	}

	return nil
}

func isValidKeyword(word []byte) bool {
	for _, keyword := range keywords {
		if bytes.Compare(word, keyword) == 0 {
			return true
		}
	}
	return false
}

func isWhiteSpace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
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

	case '[':
		if err := l.readArray(); err != nil {
			return err
		}

	case ',':
		if l.peak() != '\n' {
			return fmt.Errorf("trailing comma")
		}

	default:
		if isLetter(l.ch) {
			if err := l.readKeyword(); err != nil {
				return err
			}
		}
	}

	l.readChar()
	return nil
}
