package lexer

import (
	"mk-lang/token"
)

type charCallback func(string)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
	onChar       charCallback
}

// create lexer
func New(input string) *Lexer {
	return &Lexer{input: input, position: 0, readPosition: 0}
}

func isAlpha(char string) bool {
	return (char >= "a" && char <= "z") ||
		(char >= "A" && char <= "Z") ||
		char == "_"
}

func (l *Lexer) isAtEnd() bool {
	return l.position >= len(l.input)
}

func (l *Lexer) OnChar(callback charCallback) {
	l.onChar = callback
}

func (l *Lexer) readChar() {
	if l.position >= len(l.input) {
		l.ch = 03
	} else {
		l.ch = l.input[l.position]
	}
	l.onChar(string(l.ch))
	l.readPosition = l.position
	l.position += 1
}

func (l *Lexer) skip() {
	l.readPosition = l.position
	l.position += 1
}

func (l *Lexer) NextToken() token.Token {

	var tok token.Token

	l.skipWhitespace()
	char := l.ch
	switch char {
	case '\'':
		litteral := l.readLitteral()
		if litteral != "\\0" {
			tok = token.NewToken(token.LITTERAL, litteral)
		} else {
			tok = token.NewToken(token.ILLEGAL, "")
		}
	case '=':
		tok = token.NewToken(token.ASSIGN, string(l.ch))
	case '{':
		tok = token.NewToken(token.LBRACE, string(l.ch))
	case '}':
		tok = token.NewToken(token.RBRACE, string(l.ch))
	case '(':
		tok = token.NewToken(token.LPAREN, string(l.ch))
	case ')':
		tok = token.NewToken(token.RPAREN, string(l.ch))
	case 03:
		tok = token.NewToken(token.EOF, "")
	default:
		if isAlpha(string(char)) {
			identifier := l.readIdentifier()
			tokenType, ok := token.TokenMap[identifier]
			if ok {
				tok = token.NewToken(tokenType, identifier)
			}
			tok = token.NewToken(token.IDENT, identifier)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) readLitteral() string {
	start := l.readPosition
	for string(l.peek()) != "'" && !l.isAtEnd() {
		l.readChar()
	}
	if string(l.peek()) == "'" {
		l.readChar()
		return l.input[start:l.position]
	}
	return "\\0"
}

func (l *Lexer) skipWhitespace() {

	for l.ch == 0 || l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	// [start end [¨
	start := l.readPosition
	for isAlpha(string(l.ch)) {
		l.readChar()
	}
	return l.input[start:l.readPosition]
}

func (l *Lexer) peek() byte {
	if l.isAtEnd() {
		return 0
	}
	return l.input[l.position]
}

func (l *Lexer) ReadAll() []string {
	result := make([]string, 0)
	for l.isAtEnd() != true {
		l.readChar()
		result = append(result, string(l.ch))
	}
	return result
}

func (l *Lexer) ScanTokens() []token.Token {
	list := make([]token.Token, 0)
	for l.isAtEnd() != true {
		tok := l.NextToken()
		if tok.Type != token.SPACE {
			list = append(list, tok)
		}
	}

	return list
}
