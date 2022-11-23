package parser

import (
	"fmt"

	"github.com/lukekeum/garithmetic/compiler/parser/stack"
	"github.com/lukekeum/garithmetic/store"
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

func (p *Parser) Execute() {
	fmt.Printf("Start Parsing, (token_size=%d)\n", len(p.tokenStore))
	begin := 0

	for true {
		if begin >= len(p.tokenStore)-1 {
			break
		}

		begin, _ = p.rootExecute(begin, len(p.tokenStore)-1)

		if len(p.errorMessages) != 0 {
			panic(p.errorMessages[len(p.errorMessages)-1])
		}
	}

	fmt.Println("Finish Parsing")
}
