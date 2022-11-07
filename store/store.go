package store

type Token struct {
	TokenType TokenType
	Value     interface{}
	Line      int
}

type TokenType int

const (
	CONST TokenType = iota
	SEPER
	OPER
)

func New(tokenType TokenType, line int, value interface{}) *Token {
	return &Token{TokenType: tokenType, Line: line, Value: value}
}
