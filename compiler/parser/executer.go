package parser

import (
	"errors"
	"fmt"

	"github.com/lukekeum/garithmetic/store"
)

func (p *Parser) rootExecute(begin int, end int) (int, error) {
	t := 0
	i := begin
	var err error = nil

	for i = begin; i < end; i++ {
		begin, err = p.addExecute(i, end)
		if err != nil {
			return -1, err
		}
		i = t
		if !p.tokenStore[i].Compare(store.SEPER, ";") {
			return -1, err
		}
		i++
	}

	return i, err
}

func (p *Parser) factorExecute(begin int, end int) (int, error) {
	i := begin
	token := p.tokenStore[i]
	if token.CompareType(store.CONST) {
		p.stack.Push(token.Value.(int))
	} else if token.Compare(store.OPER, store.LPAREN) {
		try := i + 1
		bg, err := p.addExecute(try, end)
		if err != nil {
			msg := fmt.Sprintf("[Error] Expression error")
			p.errorMessages = append(p.errorMessages, msg)
			return -1, errors.New(msg)
		}
		if !(p.tokenStore[bg].Compare(store.OPER, store.RPAREN)) {
			msg := fmt.Sprintf("[Error] Unknown operator, expected ), but not found")
			p.errorMessages = append(p.errorMessages, msg)
			return -1, errors.New(msg)
		}
		i = bg + 1
	} else {
		return -1, errors.New("")
	}
	return i, nil
}

func (p *Parser) mulOper(begin int, end int) (int, error) {
	i := begin
	if !p.tokenStore[i].Compare(store.OPER, store.MULTIPLY) && !p.tokenStore[i].Compare(store.OPER, store.DIVIDE) {
		return i, nil
	}
	msg := fmt.Sprintf("[Error] Unexpected operator, expected * and /, but found %c", p.tokenStore[i].Value)
	p.errorMessages = append(p.errorMessages, msg)
	return -1, errors.New(msg)
}

func (p *Parser) mulExecute(begin int, end int) (int, error) {
	i := begin
	t, err := p.factorExecute(i, end)
	if err != nil {
		return -1, errors.New("")
	}

	for true {
		try := i
		token := p.tokenStore[try]
		t, err = p.mulOper(try, end)
		if err != nil {
			break
		}
		try += 1
		t, err = p.factorExecute(try, end)
		if err != nil {
			return -1, errors.New("")
		}
		if token.Compare(store.OPER, store.MULTIPLY) {
			p.stack.Put("mul")
		} else if token.Compare(store.OPER, store.DIVIDE) {
			p.stack.Put("div")
		}
		i = t
	}

	return i, nil
}

func (p *Parser) addExecute(begin int, end int) (int, error) {
	i := begin
	t, err := p.mulExecute(i, end)

	if err != nil {
		return -1, errors.New("")
	}

	i = t
	for true {
		try := i
		token := p.tokenStore[try]
		if !token.Compare(store.OPER, store.PLUS) && !token.Compare(store.OPER, store.MINUS) {
			break
		}
		try += 1
		t, err = p.mulExecute(try, end)
		if err != nil {
			break
		}
		if token.Compare(store.OPER, store.PLUS) {
			p.stack.Put("sub")
		} else if token.Compare(store.OPER, store.MINUS) {
			p.stack.Put("sub")
		}
	}

	return i, nil
}
