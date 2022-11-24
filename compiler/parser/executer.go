package parser

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/lukekeum/garithmetic/store"
)

// 최초로 실행되는 파서 실행 함수
func (p *Parser) rootExecute(begin int, end int) (int, error) {
	t := 0
	i := begin
	var err error = nil

	// begin부터 end까지 반복
	// 한 번 반복할 때마다 i에 1씩 더함
	for i = begin; i < end; i++ {
		t, err = p.addExecute(i, end) // 2순위 처리인 덧셈/뺄셈 감지 함수 실행
		if err != nil {               // 에러가 있으면 중지
			return -1, err
		}
		i = t // 커서(i)를 t로 변경
		// i번째 토큰이 세미콜론이면 중지
		if !p.tokenStore[i].Compare(store.SEPER, store.SEMICOLUMN) {
			p.stack.Fin()
			return -1, err
		}
		p.stack.Fin() // 구문의 끝임을 알리는 stack.Fin함수를 호출함
	}

	return i, err
}

// 정수 및 괄호 처리 함수
func (p *Parser) factorExecute(begin int, end int) (int, error) {
	i := begin
	token := p.tokenStore[i]            // 토큰 저장소에서 i번째의 토큰을 가져옴
	if token.CompareType(store.CONST) { // 토큰의 타입이 상수(정수)라면 이래를 실행함
		p.stack.Push(strconv.Itoa(token.Value)) // 토큰의 값을 문자로 변경 후 스택 저장소에 푸시함
		i++                                     // 커서를 한 칸 오른쪽으로 옮김
	} else if token.Compare(store.OPER, store.LPAREN) { // 토큰의 타입이 연산자이고 열린괄호라면 아래를 실행함
		try := i + 1                      // i + 1번째의 토큰을 테스트해보기 위해 try 변수를 상성
		bg, err := p.addExecute(try, end) // 덧셈 및 뺄셈의 구문분석을 실행하고, bg에는 이후의 토큰을 확일할 커서를, err에는 직전의 오류를 받아옴
		if err != nil {
			return -1, errors.New("") // 에러가 있으면 그대로 반환함
		}
		if !(p.tokenStore[bg].Compare(store.OPER, store.RPAREN)) { // 만약 bg번째의 토큰이 닫힌괄호가 아니면 닫힌괄호가 빠졌다고 판단후 에러 반환
			msg := fmt.Sprintf("[Error] Unknown operator, expected ), but not found")
			p.errorMessages = append(p.errorMessages, msg)
			return -1, errors.New(msg)
		}
		i = bg + 1 // i번째 토큰 커서를 닫힌괄호 이후로 설정
	} else {
		return -1, errors.New("") // 만약 i번째 토큰이 상수(정수)나 열린괄호가 아니라면 오류를 반환
	}
	return i, nil // 정수형 i와 에러는 없는 상태로 반환함
}

// begin번째의 토큰 확인하는 함수
// '*'이거나 '/'일 경우 에러 없이 1을 더해 반환, 아닐 경우 에러를 반환
func (p *Parser) mulOper(begin int, end int) (int, error) {
	i := begin
	if p.tokenStore[i].Compare(store.OPER, store.MULTIPLY) || p.tokenStore[i].Compare(store.OPER, store.DIVIDE) {
		i += 1
		return i, nil
	}
	msg := ""
	if !p.tokenStore[i].CompareType(store.OPER) && !p.tokenStore[i].CompareType(store.SEPER) {
		msg = fmt.Sprintf("[Error] Unexpected operator, expected * and /, but found %c", p.tokenStore[i].Value)
		p.errorMessages = append(p.errorMessages, msg)
	}
	return -1, errors.New(msg)
}

// 곱셈 나눗셈 처리 함수 (1순위)
func (p *Parser) mulExecute(begin int, end int) (int, error) {
	i := begin
	t, err := p.factorExecute(i, end) // i번째 토큰이 상수 또는 괄호인지 확인
	if err != nil {
		return -1, errors.New("")
	}

	for true { // 무한 반복
		try := t
		token := p.tokenStore[try]   // try번째 토큰을 불러옴
		t, err = p.mulOper(try, end) // try번째 토큰의 연산자 확인
		if err != nil {              // 에러가 있음 무한 반복을 중단함
			i = try
			break
		}
		try = t
		t, err = p.factorExecute(try, end) // try번째 토큰이 상수 또는 괄호인지 확인
		if err != nil {                    // 에러가 있음 무한 반복을 중단함
			i = try
			break
		}
		// try번째 토큰이 곱셈일 경우
		if token.Compare(store.OPER, store.MULTIPLY) {
			p.stack.Put("mul") // stack에 mul을 넣음
		}
		// try번째 토큰이 나눗셈일 경우
		if token.Compare(store.OPER, store.DIVIDE) {
			p.stack.Put("div") // stack에 div를 넣음
		}
		i = t // 커서(i)를 t로 옮긴 후 반복
	}

	return i, nil // i와 함께 nil을 반환
}

// 덧셈 뺄셈 처리 함수 (2순위)
func (p *Parser) addExecute(begin int, end int) (int, error) {
	i := begin
	t, err := p.mulExecute(i, end) // 1순위인 곱셈인지 확인함

	if err != nil {
		return -1, errors.New("")
	}

	i = t // 커서(i)를 t로 이동

	for true {
		try := i
		token := p.tokenStore[try] // token에 try번째 토큰을 담음
		// 토큰이 +이나 -인지 확인
		if !token.Compare(store.OPER, store.PLUS) && !token.Compare(store.OPER, store.MINUS) {
			break
		}
		try += 1                        // 다음 칸으로 이동
		t, err = p.mulExecute(try, end) // 1순위인 곱셈인지 확인
		nToken := p.tokenStore[t]       // t번째 토큰을 nToken에 담음
		if err != nil {                 // 에러가 있을 경우 무한반복 중지
			break
		}
		// try번째 토큰이 +인지 확인
		if token.Compare(store.OPER, store.PLUS) {
			p.stack.Put("add")
		}
		// try번째 토큰이 -인지 확인
		if token.Compare(store.OPER, store.MINUS) {
			p.stack.Put("sub")
		}

		// t번째 토큰이 세미콜론이나 닫힌 괄호라면 무한반복 중지
		if nToken.Compare(store.SEPER, store.SEMICOLUMN) || nToken.Compare(store.OPER, store.RPAREN) {
			i = t
			break
		}
	}

	return i, nil
}
