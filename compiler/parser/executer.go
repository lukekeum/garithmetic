package parser

import (
	"errors"

	"github.com/lukekeum/garithmetic/store"
)

func (p *Parser) rootExecute(begin int, end int, tstore []store.Token) (int, error) {
	t := 0
	err := errors.New("")

	for i := begin; i < end; i++ {
		begin, err = p.addExecute(i, end, tstore)
		if err != nil {
			return begin, err
		}
		i = t
	}

	return begin, nil
}

func (p *Parser) addExecute(begin int, end int, tstore []store.Token) (int, error) {
	return begin, nil
}
