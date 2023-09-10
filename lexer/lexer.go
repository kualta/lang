package lexer

import (
	"kulang/token"
	"strings"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isIdentifier(l.ch) {
			tok.Literal, tok.Type = l.readToken()
			return tok
		}
	}
	l.readChar()

	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readToken() (string, token.TokenType) {
	startPosition := l.position

	for isIdentifier(l.ch) {
		l.readChar()
	}

	tokenIdent := l.input[startPosition:l.position]
	tokenType := getTokenType(tokenIdent)

	return tokenIdent, tokenType
}

func getTokenType(tokenIdent string) token.TokenType {
	var tokenType token.TokenType = token.ILLEGAL

	if strings.ContainsRune(tokenIdent, '.') {
		tokenType = token.FLOAT
	} else if isDigit(tokenIdent[len(tokenIdent)-1]) {
		tokenType = token.INT
	} else if isLetter(tokenIdent[len(tokenIdent)-1]) {
		tokenType = token.LookupIdent(tokenIdent)
	}

	return tokenType
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	//                               \/ allow for 1_000_000
	return '0' <= ch && ch <= '9' || ch == '_' || ch == '.'
}

func isIdentifier(ch byte) bool {
	return isLetter(ch) || isDigit(ch)
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}
