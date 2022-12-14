package store

import (
	"fmt"
)

// 토큰 객체에 대한 정의
type Token struct {
	TokenType   TokenType
	Value       int
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

func GetTokenKeyFromInteger(tokenNum TokenType) string {
	switch tokenNum {
	case 0:
		return "CONST"
	case 1:
		return "SEPERATOR"
	case 2:
		return "OPERATOR"
	default:
		panic("[Error] Unexcepted Token : expected 'CONST' or 'SEPERATOR' or 'OPERATOR' but found nil")
	}
}

func IntToTokenNumber(n int) TokenNumber {
	switch n {
	case 43:
		return PLUS
	case 45:
		return MINUS
	case 42:
		return MULTIPLY
	case 47:
		return DIVIDE
	case 37:
		return 37
	case 40:
		return LPAREN
	case 41:
		return RPAREN
	case 59:
		return SEMICOLUMN
	default:
		return SEMICOLUMN
	}
}

func (t TokenNumber) Number() int {
	switch t {
	case PLUS:
		return 43
	case MINUS:
		return 45
	case MULTIPLY:
		return 42
	case DIVIDE:
		return 47
	case REMAINDER:
		return 37
	case LPAREN:
		return 40
	case RPAREN:
		return 41
	case SEMICOLUMN:
		return 59
	default:
		return -1
	}
}

// 토큰 객체 생성 (토큰 저장 시 이용)
func New(tokenType TokenType, line int, parsingLine int, value int) *Token {
	switch tokenType {
	case CONST: // 저장하려는 토큰의 타입이 상수일 경우
		return &Token{TokenType: tokenType, Line: line, ParsingLine: parsingLine, Value: value}
	case SEPER: // 저장하려는 토큰의 타입이 구분자일 경우
		if value == ';' {
			return &Token{TokenType: tokenType, Line: line, ParsingLine: parsingLine, Value: 59}
		} else { // 정의된 구분자 ';'가 아닌 경우 에러 방출
			panic(fmt.Sprintf("[Error] Unexpected Token : expected ';' but found %c", value))
		}
	case OPER: // 저장하려는 토큰의 타입이 연산자일 경우
		return checkOperValue(value, line, parsingLine)
	}

	return &Token{TokenType: tokenType, Line: line, ParsingLine: parsingLine, Value: value}
}

func checkOperValue(value interface{}, line int, parsingLine int) *Token {
	switch fmt.Sprintf("%v", value) {
	case "43": // 저장하려는 토큰이 + 일 경우
		return &Token{TokenType: OPER, Line: line, ParsingLine: parsingLine, Value: 43}
	case "45": // 저장하려는 토큰이 - 일 경우
		return &Token{TokenType: OPER, Line: line, ParsingLine: parsingLine, Value: 45}
	case "42": // 저장하려는 토큰이 * 일 경우
		return &Token{TokenType: OPER, Line: line, ParsingLine: parsingLine, Value: 42}
	case "47": // 저장하려는 토큰이 / 일 경우
		return &Token{TokenType: OPER, Line: line, ParsingLine: parsingLine, Value: 47}
	case "37": // 저장하려는 토큰이 % 일 경우
		return &Token{TokenType: OPER, Line: line, ParsingLine: parsingLine, Value: 37}
	case "40": // 저장하려는 토큰이 ( 일 경우
		return &Token{TokenType: OPER, Line: line, ParsingLine: parsingLine, Value: 40}
	case "41": // 저장하려는 토큰이 ) 일 경우
		return &Token{TokenType: OPER, Line: line, ParsingLine: parsingLine, Value: 41}
	default: // 정의된 연산자 '+, -, *, /, %, (, )'가 아닌 경우 에러 방출
		panic(fmt.Sprintf("[Error] Unexpected Token : expected operator but found %c", value))
	}
}

func (t *Token) Compare(tokenType TokenType, value TokenNumber) bool {
	return t.TokenType == tokenType && t.Value == value.Number()
}

func (t *Token) CompareType(tokenType TokenType) bool {
	return t.TokenType == tokenType
}
