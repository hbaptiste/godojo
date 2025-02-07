package token

import (
	"fmt"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var TokenMap = map[string]TokenType{
	"func":   FUNCTION,
	"let":    LET,
	"return": RETURN,
	"const":  CONST,
}

var TokenToNameMap = map[string]TokenType{
	",": "COMMA",
	";": "SEMICOLON",
	"(": "LPAREN",
	")": "RPAREN",
	"{": "LBRACE",
	"}": "RBRACE",
	"=": "EQUAL",
	"return" : "RETURN",
}

func (token Token) String() string {
	return fmt.Sprintf("<Token type='%v' value=%v />", TokenToNameMap[string(token.Type)], token.Literal)
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	SPACE   = "SPACE"

	// Identifiers + literals
	IDENT    = "IDENT"
	INT      = "INT"
	LITTERAL = "LITTERAL"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimeters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNC"
	LET      = "LET"
	RETURN   = "RETURN"
	CONST    = "CONST"
)

func NewToken(tokenType TokenType, value string) Token {
	return Token{tokenType, value}
}
