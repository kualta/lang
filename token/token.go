package token

import "strings"

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if token, ok := keywords[ident]; ok {
		return token
	}
	return IDENT
}

func GetTokenType(tokenIdent string) TokenType {
	var tokenType TokenType = ILLEGAL

	if strings.ContainsRune(tokenIdent, '.') {
		tokenType = FLOAT
	} else if IsDigit(tokenIdent[len(tokenIdent)-1]) {
		tokenType = INT
	} else if IsLetter(tokenIdent[len(tokenIdent)-1]) {
		tokenType = LookupIdent(tokenIdent)
	}

	return tokenType
}

func IsLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func IsDigit(ch byte) bool {
	//                               \/ allow for 1_000_000
	return '0' <= ch && ch <= '9' || ch == '_' || ch == '.'
}

func TrimDigit(ident string) string {
	return strings.ReplaceAll(ident, "_", "")
}

func IsIdentifier(ch byte) bool {
	return IsLetter(ch) || IsDigit(ch)
}

const (
	// Special
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT" // x, y
	FLOAT = "FLOAT" // float numbers
	INT   = "INT"   // 0 - 9

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"
	EQ = "=="
	NE = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)
