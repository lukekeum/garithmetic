package store

import "fmt"

type Token struct {
	TokenType   TokenType
	Value       interface{}
	Line        int
	ParsingLine int
}

type TokenType int
type TokenNumber int

const (
	CONST TokenType = iota
	SEPER
	OPER
)

const (
	PLUS       TokenNumber = 43
	MINUS      TokenNumber = 45
	MULTIPLY   TokenNumber = 42
	DIVIDE     TokenNumber = 47
	REMAINDER  TokenNumber = 37
	LPAREN     TokenNumber = 40
	RPAREN     TokenNumber = 41
	SEMICOLUMN TokenNumber = 59
)

func New(tokenType TokenType, line int, parsingLine int, value interface{}) *Token {
	switch tokenType {
	case CONST:
		return &Token{TokenType: tokenType, Line: line, ParsingLine: parsingLine, Value: value}
	case SEPER:
		if value == ';' {
			return &Token{TokenType: tokenType, Line: line, ParsingLine: parsingLine, Value: 59}
		} else {
			panic(fmt.Sprintf("[Error] Unexpected Token : expected ';' but found %c", value))
		}
	case OPER:
		switch value {
		case '+':
			return &Token{TokenType: tokenType, Line: line, ParsingLine: parsingLine, Value: 43}
		case '-':
			return &Token{TokenType: tokenType, Line: line, ParsingLine: parsingLine, Value: 45}
		case '*':
			return &Token{TokenType: tokenType, Line: line, ParsingLine: parsingLine, Value: 42}
		case '/':
			return &Token{TokenType: tokenType, Line: line, ParsingLine: parsingLine, Value: 47}
		case '%':
			return &Token{TokenType: tokenType, Line: line, ParsingLine: parsingLine, Value: 37}
		case '(':
			return &Token{TokenType: tokenType, Line: line, ParsingLine: parsingLine, Value: 40}
		case ')':
			return &Token{TokenType: tokenType, Line: line, ParsingLine: parsingLine, Value: 41}
		default:
			panic(fmt.Sprintf("[Error] Unexpected Token : expected operator but found %c", value))
		}
	}

	return &Token{TokenType: tokenType, Line: line, ParsingLine: parsingLine, Value: value}
}
