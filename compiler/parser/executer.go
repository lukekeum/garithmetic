package parser

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/lukekeum/garithmetic/store"
)

func (p *Parser) rootExecute(begin int, end int) (int, error) {
	t := 0
	i := begin
	var err error = nil

	for i = begin; i < end; i++ {
		t, err = p.addExecute(i, end)
		if err != nil {
			return -1, err
		}
		i = t
		if !p.tokenStore[i].Compare(store.SEPER, store.SEMICOLUMN) {
			p.stack.Fin()
			return -1, err
		}
	}

	return i, err
}

func (p *Parser) factorExecute(begin int, end int) (int, error) {
	i := begin
	token := p.tokenStore[i]
	if token.CompareType(store.CONST) {
		p.stack.Push(strconv.Itoa(token.Value))
		i++
	} else if token.Compare(store.OPER, store.LPAREN) {
		try := i + 1
		bg, err := p.addExecute(try, end)
		if err != nil {
			return -1, errors.New("")
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

func (p *Parser) mulExecute(begin int, end int) (int, error) {
	i := begin
	t, err := p.factorExecute(i, end)
	if err != nil {
		return -1, errors.New("")
	}

	for true {
		try := t
		token := p.tokenStore[try]
		t, err = p.mulOper(try, end)
		if err != nil {
			i = try
			break
		}
		try = t
		t, err = p.factorExecute(try, end)
		if err != nil {
			i = try
			break
		}
		if token.Compare(store.OPER, store.MULTIPLY) {
			p.stack.Put("mul")
		}
		if token.Compare(store.OPER, store.DIVIDE) {
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
		nToken := p.tokenStore[t]
		if err != nil {
			break
		}
		if token.Compare(store.OPER, store.PLUS) {
			p.stack.Put("add")
		} else if token.Compare(store.OPER, store.MINUS) {
			p.stack.Put("sub")
		}

		if nToken.Compare(store.SEPER, store.SEMICOLUMN) || nToken.Compare(store.OPER, store.RPAREN) {
			i = t
			break
		}
	}

	return i, nil
}
