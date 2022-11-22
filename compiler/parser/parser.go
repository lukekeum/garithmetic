package parser

import (
	"github.com/lukekeum/garithmetic/compiler/parser/stack"
	"github.com/lukekeum/garithmetic/store"
	"github.com/lukekeum/garithmetic/tree"
)

type Parser struct {
	tokenStore    []store.Token
	stack         stack.Stack
	errorMessages []string
}

func New(store []store.Token) *Parser {
	stack := stack.Stack{}
	return &Parser{tokenStore: store, stack: stack}
}

func (p *Parser) Execute() []*tree.Tree {
	tstore := p.tokenStore
	tree := []*tree.Tree{}
	begin := 0
	var err interface{} = nil

	for i := 0; i < len(tstore); i++ {
		token := tstore[i]
		if token.Compare(store.SEPER, store.SEMICOLUMN) {
			begin, err = p.rootExecute(begin, i)
			if err != nil {
				panic(err)
			}
			begin = i
		}
	}

	return tree
}
