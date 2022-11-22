package parser

import (
	"github.com/lukekeum/garithmetic/compiler/parser/stack"
	"github.com/lukekeum/garithmetic/store"
	"github.com/lukekeum/garithmetic/tree"
)

type Parser struct {
	tokenStore []store.Token
	stack      stack.Stack
}

func New(store []store.Token) *Parser {
	stack := stack.Stack{}
	return &Parser{tokenStore: store, stack: stack}
}

func (p *Parser) Execute() []*tree.Tree {
	tstore := p.tokenStore
	tree := []*tree.Tree{}
	begin := 0

	for i := 0; i < len(tstore); i++ {
		token := tstore[i]
		if token.Compare(store.SEPER, store.SEMICOLUMN) {
			t, err := rootExecute(begin, i, tstore)
			if err != nil {
				panic(err)
			}
			tree = append(tree, t)
			begin = i
		}
	}

	return tree
}
